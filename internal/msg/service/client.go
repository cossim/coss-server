package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cossim/coss-server/internal/msg/api/http/model"
	relationgrpcv1 "github.com/cossim/coss-server/internal/relation/api/grpc/v1"
	"github.com/cossim/coss-server/pkg/cache"
	"github.com/cossim/coss-server/pkg/constants"
	"github.com/cossim/coss-server/pkg/msg_queue"
	pkgtime "github.com/cossim/coss-server/pkg/utils/time"
	"github.com/gorilla/websocket"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"strconv"
	"sync"
)

type client struct {
	Conn           *websocket.Conn
	Uid            string //客户端所有者
	Rid            int64  //客户端id
	ClientType     string //客户端类型
	DriverId       string //客户端id
	queue          *amqp091.Channel
	wsMutex        sync.Mutex
	redisMutex     sync.Mutex
	Rdb            *cache.RedisClient
	relationClient relationgrpcv1.UserRelationServiceClient
}

// 用户上线
func (c *client) wsOnlineClients() {
	wsMutex.Lock()
	pool[c.Uid][c.ClientType] = append(pool[c.Uid][c.ClientType], c)
	wsMutex.Unlock()

	//通知前端接收离线消息
	//TODO 添加上线的设备类型
	msg := constants.WsMsg{Uid: c.Uid, DriverId: c.DriverId, Event: constants.OnlineEvent, Rid: c.Rid, SendAt: pkgtime.Now(), Data: constants.OnlineEventData{DriverType: constants.DriverType(c.ClientType)}}
	js, _ := json.Marshal(msg)
	if Enc == nil {
		log.Println("加密客户端错误", zap.Error(nil))
		return
	}
	message, err := Enc.GetSecretMessage(string(js), c.Uid)
	if err != nil {
		fmt.Println("加密失败：", err)
		return
	}

	//上线推送消息
	err = c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("发送消息失败：", err)
		return
	}

	err = c.pushAllFriendOnlineStatus()
	if err != nil {
		fmt.Println("推送好友在线状态失败：", err)
		return
	}

	//修改在线状态
	err = c.addUserWsCount()
	if err != nil {
		fmt.Println("修改在线状态失败：", err)
		return
	}

	for {
		msg, ok, err := msg_queue.ConsumeMessages(c.Uid, c.queue)
		if err != nil || !ok {
			//c.queue.Stop()
			//拉取完之后删除队列
			_ = rabbitMQClient.DeleteEmptyQueue(c.Uid)
			return
		}

		err = c.Conn.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			fmt.Println("发送消息失败：", err)
			return
		}
	}
}

func (c *client) wsOfflineClients() {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	if _, ok := pool[c.Uid][c.ClientType]; ok {
		for i, c2 := range pool[c.Uid][c.ClientType] {
			if c2.Rid == c.Rid {
				pool[c.Uid][c.ClientType] = append(pool[c.Uid][c.ClientType][:i], pool[c.Uid][c.ClientType][i+1:]...)
				if len(pool[c.Uid][c.ClientType]) == 0 {
					delete(pool[c.Uid], c.ClientType)
				}
				break
			}
		}
	}
	err := c.reduceUserWsCount()
	if err != nil {
		fmt.Println("修改在线状态失败：", err)
		return
	}
}

func (c *client) addUserWsCount() error {
	c.redisMutex.Lock()
	defer c.redisMutex.Unlock()
	exists, err := c.Rdb.ExistsKey(c.Uid)
	if err != nil {
		return err
	}

	//给好友推送上线
	err = c.pushFriendStatus(onlineEvent)
	if err != nil {
		return err
	}

	if !exists {
		return c.Rdb.SetKey(c.Uid, 1, 0)
	} else {
		value, err := c.Rdb.GetKey(c.Uid)
		if err != nil {
			return err
		}
		str := value.(string)
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		num++
		return c.Rdb.SetKey(c.Uid, num, 0)
	}
}

func (c *client) reduceUserWsCount() error {
	c.redisMutex.Lock()
	defer c.redisMutex.Unlock()
	exists, err := c.Rdb.ExistsKey(c.Uid)
	if err != nil {
		return err
	}
	if !exists {
		//给好友推送下线
		err := c.pushFriendStatus(offlineEvent)
		if err != nil {
			return err
		}
		return nil
	} else {
		value, err := c.Rdb.GetKey(c.Uid)
		if err != nil {
			return err
		}
		str := value.(string)
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		fmt.Printf("%s账号当前还有%d个客户端在线", c.Uid, num)
		if num == 1 {
			//给好友推送下线
			err := c.pushFriendStatus(offlineEvent)
			if err != nil {
				return err
			}
			return c.Rdb.DelKey(c.Uid)
		} else {
			num--
			return c.Rdb.SetKey(c.Uid, num, 0)
		}
	}
}

type status uint

const (
	onlineEvent status = iota + 1
	// OfflineEvent 下线事件
	offlineEvent
)

// 给好友推送离线或上线通知
func (c *client) pushFriendStatus(s status) error {
	//查询所有好友
	list, err := c.relationClient.GetFriendList(context.Background(), &relationgrpcv1.GetFriendListRequest{UserId: c.Uid})
	if err != nil {
		return err
	}
	if len(list.FriendList) > 0 {
		for _, friend := range list.FriendList {
			msg := constants.WsMsg{Uid: friend.UserId, Event: constants.FriendUpdateOnlineStatusEvent, Rid: c.Rid, SendAt: pkgtime.Now(), Data: model.FriendOnlineStatusMsg{Status: int32(s), UserId: c.Uid}}
			js, _ := json.Marshal(msg)
			if Enc == nil {
				log.Println("加密客户端错误", zap.Error(nil))
				continue
			}
			message, err := Enc.GetSecretMessage(string(js), c.Uid)
			if err != nil {
				fmt.Println("加密失败：", err)
				continue
			}

			for _, v := range pool[friend.UserId] {
				for _, cc := range v {
					if cc.Conn != nil && cc.Conn.WriteMessage(websocket.PingMessage, nil) != websocket.ErrCloseSent {
						err := cc.Conn.WriteMessage(websocket.TextMessage, []byte(message))
						if err != nil {
							log.Println("发送消息失败：", err)
							break
						}
					}

				}
			}
		}

	}
	return nil
}

// 获取所有好友在线状态
func (c *client) pushAllFriendOnlineStatus() error {
	fmt.Println("pushAllFriendOnlineStatus", c.relationClient)
	if c.Conn == nil || c.Conn.WriteMessage(websocket.PingMessage, nil) != nil {
		log.Println("连接已关闭，不再推送消息")
		return nil
	}
	//查询所有好友
	list, err := c.relationClient.GetFriendList(context.Background(), &relationgrpcv1.GetFriendListRequest{UserId: c.Uid})
	if err != nil {
		return err
	}
	var friendList []model.FriendOnlineStatusMsg

	if len(list.FriendList) > 0 {
		for _, friend := range list.FriendList {
			exists, err := c.Rdb.ExistsKey(friend.UserId)
			if err != nil {
				return err
			}
			if exists {
				value, err := c.Rdb.GetKey(friend.UserId)
				if err != nil {
					return err
				}
				str := value.(string)
				num, err := strconv.Atoi(str)
				if err != nil {
					return err
				}
				if num > 0 {
					friendList = append(friendList, model.FriendOnlineStatusMsg{Status: int32(onlineEvent), UserId: friend.UserId})
				} else {
					friendList = append(friendList, model.FriendOnlineStatusMsg{Status: int32(offlineEvent), UserId: friend.UserId})
				}
			} else {
				friendList = append(friendList, model.FriendOnlineStatusMsg{Status: int32(onlineEvent), UserId: friend.UserId})
			}
		}
	}

	msg := constants.WsMsg{Uid: c.Uid, Event: constants.PushAllFriendOnlineStatusEvent, Rid: c.Rid, SendAt: pkgtime.Now(), Data: friendList}
	js, _ := json.Marshal(msg)
	if Enc == nil {
		log.Println("加密客户端错误", zap.Error(nil))
		return nil
	}
	message, err := Enc.GetSecretMessage(string(js), c.Uid)
	if err != nil {
		fmt.Println("加密失败：", err)
		return nil
	}
	err = c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
