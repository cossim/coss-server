package service

import (
	"context"
	"fmt"
	"github.com/cossim/coss-server/interface/msg/config"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/discovery"
	"github.com/cossim/coss-server/pkg/encryption"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/msg_queue"
	groupgrpcv1 "github.com/cossim/coss-server/service/group/api/v1"
	msggrpcv1 "github.com/cossim/coss-server/service/msg/api/v1"
	relationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1"
	usergrpcv1 "github.com/cossim/coss-server/service/user/api/v1"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sync"
)

var (
	rabbitMQClient *msg_queue.RabbitMQ

	wsRid   int64 = 0 //全局客户端id
	wsMutex       = sync.Mutex{}
	Enc     encryption.Encryptor
	pool    = make(map[string]map[string][]*client)
)

// Service struct
type Service struct {
	relationUserClient   relationgrpcv1.UserRelationServiceClient
	relationGroupClient  relationgrpcv1.GroupRelationServiceClient
	relationDialogClient relationgrpcv1.DialogServiceClient
	userClient           usergrpcv1.UserServiceClient
	groupClient          groupgrpcv1.GroupServiceClient
	msgClient            msggrpcv1.MsgServiceClient
	groupMsgClient       msggrpcv1.GroupMessageServiceClient
	redisClient          *redis.Client

	//mqClient *msg_queue.RabbitMQ

	logger    *zap.Logger
	sid       string
	discovery discovery.Registry
	ac        *pkgconfig.AppConfig
	//Enc       encryption.Encryptor

	//pool  map[string]map[string][]*client
}

func New(ac *pkgconfig.AppConfig) *Service {
	mqClient, err := msg_queue.NewRabbitMQ(fmt.Sprintf("amqp://%s:%s@%s", ac.MessageQueue.Username, ac.MessageQueue.Password, ac.MessageQueue.Addr()))
	if err != nil {
		panic(err)
	}
	rabbitMQClient = mqClient

	setupEncryption(ac)

	s := &Service{
		ac:          ac,
		logger:      plog.NewDevLogger("msg_bff"),
		sid:         xid.New().String(),
		redisClient: setupRedis(ac),
		//rabbitMQClient: mqClient,
		//pool:     make(map[string]map[string][]*client),
		//mqClient: mqClient,
		//pool:     make(map[string]map[string][]*client),
	}
	// 监听服务消息队列
	go s.ListenQueue()

	return s
}

func (s *Service) ListenQueue() {
	if rabbitMQClient.GetConnection().IsClosed() {
		panic("mqClient Connection is closed")
	}
	msgs, err := rabbitMQClient.ConsumeServiceMessages(msg_queue.MsgService, msg_queue.Service_Exchange)
	if err != nil {
		panic(err)
	}
	go func() {
		//监听队列
		for msg := range msgs {
			var msg_query msg_queue.ServiceQueueMsg
			err := json.Unmarshal(msg.Body, &msg_query)
			if err != nil {
				fmt.Println("解析消息失败：", err)
				return
			}

			mmap, ok := msg_query.Data.(map[string]interface{})
			if !ok {
				fmt.Println("解析消息失败：")
				return
			}
			//map解析成结构体
			jsonData, err := json.Marshal(mmap)
			if err != nil {
				fmt.Println("Failed to marshal map to JSON:", err)
				return
			}
			var wsm config.WsMsg
			err = json.Unmarshal(jsonData, &wsm)
			if err != nil {
				fmt.Println("解析消息失败：", err)
				return
			}

			switch msg_query.Action {
			//推送消息
			case msg_queue.SendMessage:
				s.SendMsg(wsm.Uid, wsm.DriverId, wsm.Event, wsm.Data, true)
			case msg_queue.LiveEvent:
				s.SendMsg(wsm.Uid, wsm.DriverId, wsm.Event, wsm.Data, false)
			//强制断开ws
			case msg_queue.UserWebsocketClose:
				thismap, ok := wsm.Data.(map[string]interface{})
				if !ok {
					fmt.Println("解析消息失败：")
					return
				}
				t := thismap["driver_type"]
				id := thismap["rid"]
				//类型断言
				driType, ok := t.(string)
				if !ok {
					fmt.Println("解析消息失败：")
					return
				}
				rid := id.(float64)
				for _, c := range pool[wsm.Uid][driType] {

					if c.Rid == int64(rid) {
						fmt.Println("关闭连接", rid)
						c.Conn.Close()
					}
				}
			}
		}
	}()
}

func (s *Service) Stop(ctx context.Context) error {
	//关闭队列
	rabbitMQClient.Close()
	return nil
}

func (s *Service) HandlerGrpcClient(serviceName string, conn *grpc.ClientConn) error {
	fmt.Println("9999999999999999")
	switch serviceName {
	case "user_service":
		s.userClient = usergrpcv1.NewUserServiceClient(conn)
		s.logger.Info("gRPC client for user service initialized", zap.String("service", "user"), zap.String("addr", conn.Target()))
	case "relation_service":
		s.relationUserClient = relationgrpcv1.NewUserRelationServiceClient(conn)
		s.logger.Info("gRPC client for relation service initialized", zap.String("service", "userRelation"), zap.String("addr", conn.Target()))

		s.relationGroupClient = relationgrpcv1.NewGroupRelationServiceClient(conn)
		s.logger.Info("gRPC client for relation service initialized", zap.String("service", "groupRelation"), zap.String("addr", conn.Target()))

		s.relationDialogClient = relationgrpcv1.NewDialogServiceClient(conn)
		s.logger.Info("gRPC client for relation service initialized", zap.String("service", "dialogRelation"), zap.String("addr", conn.Target()))
	case "group_service":
		s.groupClient = groupgrpcv1.NewGroupServiceClient(conn)
		s.logger.Info("gRPC client for group service initialized", zap.String("service", "group"), zap.String("addr", conn.Target()))
	case "msg_service":
		s.groupMsgClient = msggrpcv1.NewGroupMessageServiceClient(conn)
		s.msgClient = msggrpcv1.NewMsgServiceClient(conn)
		s.logger.Info("gRPC client for group service initialized", zap.String("service", "msg"), zap.String("addr", conn.Target()))
	}

	return nil
}

func setupRedis(ac *pkgconfig.AppConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ac.Redis.Addr(),
		Password: ac.Redis.Password, // no password set
		DB:       0,                 // use default DB
		//Protocol: cfg,
	})
}

func setupEncryption(ac *pkgconfig.AppConfig) {
	enc := encryption.NewEncryptor([]byte(ac.Encryption.Passphrase), ac.Encryption.Name, ac.Encryption.Email, ac.Encryption.RsaBits, ac.Encryption.Enable)

	err := enc.ReadKeyPair()
	if err != nil {
		return
	}

	Enc = enc
}
