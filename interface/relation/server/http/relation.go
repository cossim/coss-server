package http

import (
	"context"
	"fmt"
	msgconfig "github.com/cossim/coss-server/interface/msg/config"
	"github.com/cossim/coss-server/interface/relation/api/model"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/http"
	pkghttp "github.com/cossim/coss-server/pkg/http"
	"github.com/cossim/coss-server/pkg/msg_queue"
	groupApi "github.com/cossim/coss-server/service/group/api/v1"
	"strconv"
	"time"

	"github.com/cossim/coss-server/pkg/http/response"
	"github.com/cossim/coss-server/pkg/utils/usersorter"
	relationApi "github.com/cossim/coss-server/service/relation/api/v1"
	userApi "github.com/cossim/coss-server/service/user/api/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 黑名单
// @Description 黑名单
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/user/blacklist [get]
func blackList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	// 获取黑名单列表
	blacklistResp, err := userRelationClient.GetBlacklist(context.Background(), &relationApi.GetBlacklistRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	var users []string
	for _, user := range blacklistResp.Blacklist {
		users = append(users, user.UserId)
	}

	blacklist, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: users})
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取黑名单列表成功", blacklist)
}

// @Summary 好友列表
// @Description 好友列表
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/user/friend_list [get]
func friendList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}
	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		logger.Error("user service UserInfo", zap.Error(err))
		c.Error(err)
		return
	}

	// 获取好友列表
	friendListResp, err := userRelationClient.GetFriendList(context.Background(), &relationApi.GetFriendListRequest{UserId: userID})
	if err != nil {
		logger.Error("user service GetFriendList", zap.Error(err))
		c.Error(err)
		return
	}

	var users []string
	for _, user := range friendListResp.FriendList {
		users = append(users, user.UserId)
	}

	userInfos, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: users})
	if err != nil {
		logger.Error("user service GetBatchUserInfo", zap.Error(err))
		c.Error(err)
		return
	}

	var data []usersorter.User
	for _, v := range userInfos.Users {
		for _, friend := range friendListResp.FriendList {
			if friend.UserId == v.UserId {
				data = append(data, usersorter.CustomUserData{
					UserID:   v.UserId,
					NickName: v.NickName,
					Email:    v.Email,
					Tel:      v.Tel,
					Avatar:   v.Avatar,
					Status:   uint(v.Status),
					DialogId: friend.DialogId,
				})
				break
			}
		}

	}

	// Sort and group by specified field
	groupedUsers := usersorter.SortAndGroupUsers(data, "NickName")

	response.Success(c, "获取好友列表成功", groupedUsers)
}

// @Summary 好友申请列表
// @Description 好友申请列表
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/user/request_list [get]
func userRequestList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}
	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		logger.Error("Failed to get user information", zap.Error(err))
		c.Error(err)
		return
	}

	reqList, err := userRelationClient.GetFriendRequestList(context.Background(), &relationApi.GetFriendRequestListRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	var ids []string
	var data []*model.RequestListResponse
	for _, v := range reqList.FriendRequestList {
		ids = append(ids, v.UserId)
		data = append(data, &model.RequestListResponse{
			UserID: v.UserId,
			Msg:    v.Msg,
			Status: uint32(v.Status),
		})
	}

	users, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: ids})
	if err != nil {
		c.Error(err)
		return
	}

	for _, v := range data {
		for _, u := range users.Users {
			if v.UserID == u.UserId {
				v.Nickname = u.NickName
				v.Avatar = u.Avatar
				break
			}
		}
	}

	response.Success(c, "获取好友申请列表成功", data)
}

// @Summary 删除黑名单
// @Description 删除黑名单
// @Accept  json
// @Produce  json
// @param request body model.DeleteBlacklistRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/delete_blacklist [post]
func deleteBlacklist(c *gin.Context) {
	req := new(model.DeleteBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	// 检查要删除的黑名单用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	// 进行删除黑名单操作
	if _, err = userRelationClient.DeleteBlacklist(context.Background(), &relationApi.DeleteBlacklistRequest{UserId: userID, FriendId: req.UserID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "删除黑名单成功", nil)
}

// @Summary 添加黑名单
// @Description 添加黑名单
// @Accept  json
// @Produce  json
// @param request body addBlacklistRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/add_blacklist [post]
func addBlacklist(c *gin.Context) {
	req := new(model.AddBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	// 检查添加黑名单的用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	// 进行添加黑名单操作
	if _, err = userRelationClient.AddBlacklist(context.Background(), &relationApi.AddBlacklistRequest{UserId: userID, FriendId: req.UserID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "添加到黑名单成功", nil)
}

// @Summary 删除好友
// @Description 删除好友
// @Accept  json
// @Produce  json
// @param request body model.DeleteFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/delete_friend [post]
func deleteFriend(c *gin.Context) {
	req := new(model.DeleteFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	// 检查删除的用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	relation, err := userRelationClient.GetUserRelation(context.Background(), &relationApi.GetUserRelationRequest{UserId: userID, FriendId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}
	//
	if relation.Status != relationApi.RelationStatus_RELATION_STATUS_ADDED {
		response.Fail(c, "不是好友", nil)
		return
	}
	//删除自己的对话用户
	_, err = dialogClient.DeleteDialogUserByDialogIDAndUserID(context.Background(), &relationApi.DeleteDialogUserByDialogIDAndUserIDRequest{DialogId: relation.DialogId, UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	// 进行删除好友操作
	if _, err = userRelationClient.DeleteFriend(context.Background(), &relationApi.DeleteFriendRequest{UserId: userID, FriendId: req.UserID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "删除好友成功", nil)
}

// @Summary 管理好友请求
// @Description 管理好友请求
// @Accept  json
// @Produce  json
// @param request body model.ManageFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/manage_friend [post]
func manageFriend(c *gin.Context) {
	req := new(model.ManageFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	// 检查要添加的用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	var status relationApi.RelationStatus
	var dialogId uint32 = 0
	if req.Status == 1 {
		status = relationApi.RelationStatus_RELATION_STATUS_ADDED
		relation, err := userRelationClient.GetUserRelation(context.Background(), &relationApi.GetUserRelationRequest{UserId: userID, FriendId: req.UserID})
		if err != nil {
			c.Error(err)
			return
		}
		if relation != nil && relation.DialogId != 0 {
			fmt.Println("之前已经有关系，直接加入对话")
			_, err = dialogClient.JoinDialog(context.Background(), &relationApi.JoinDialogRequest{DialogId: relation.DialogId, UserId: req.UserID})
			if err != nil {
				c.Error(err)
				return
			}
			dialogId = relation.DialogId
		} else {
			//创建对话
			dialog, err := dialogClient.CreateDialog(context.Background(), &relationApi.CreateDialogRequest{OwnerId: userID, Type: 0, GroupId: 0})
			if err != nil {
				c.Error(err)
				return
			}
			//加入对话
			_, err = dialogClient.JoinDialog(context.Background(), &relationApi.JoinDialogRequest{DialogId: dialog.Id, UserId: userID})
			if err != nil {
				c.Error(err)
				return
			}
			//确认添加好友之后加入对话
			_, err = dialogClient.JoinDialog(context.Background(), &relationApi.JoinDialogRequest{DialogId: dialog.Id, UserId: req.UserID})
			if err != nil {
				c.Error(err)
				return
			}
			dialogId = dialog.Id
		}
	} else {
		status = relationApi.RelationStatus_RELATION_STATUS_REJECTED
	}
	// 进行确认好友操作
	if _, err = userRelationClient.ManageFriend(context.Background(), &relationApi.ManageFriendRequest{UserId: userID, FriendId: req.UserID, DialogId: dialogId, Status: status}); err != nil {
		c.Error(err)
		return
	}
	targetId := req.UserID
	req.UserID = userID
	targetInfo, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: targetId})
	if err != nil {
		return
	}
	myInfo, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		return
	}
	wsMsgData := map[string]interface{}{"user_id": userID, "status": req.Status}
	msg := msgconfig.WsMsg{Uid: targetId, Event: msgconfig.ManageFriendEvent, Data: wsMsgData}
	var responseData interface{}
	if req.Status == 1 {
		wsMsgData["target_info"] = myInfo
		wsMsgData["e2e_public_key"] = req.E2EPublicKey
		responseData = targetInfo
	}
	err = rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
	if err != nil {
		logger.Error("推送服务消息失败", zap.Error(err))
		response.Fail(c, "管理好友申请失败", nil)
		return
	}

	response.Success(c, "管理好友申请成功", responseData)
}

// @Summary 添加好友
// @Description 添加好友
// @Accept  json
// @Produce  json
// @param request body model.AddFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/add_friend [post]
func addFriend(c *gin.Context) {
	req := new(model.AddFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}
	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: thisId})
	if err != nil {
		//logger.Error("user service", zap.Error(err))
		//response.Fail(c, "用户不存在", nil)
		err = code.UserErrNotExist
		c.Error(err)
		return
	}
	// 检查添加的用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		err = code.UserErrNotExist
		c.Error(err)
		return
	}

	if _, err = userRelationClient.AddFriend(context.Background(), &relationApi.AddFriendRequest{
		UserId:   thisId,
		FriendId: req.UserID,
		Msg:      req.Msg,
	}); err != nil {
		c.Error(err)
		return
	}
	targetId := req.UserID
	req.UserID = thisId
	msg := msgconfig.WsMsg{Uid: targetId, Event: msgconfig.AddFriendEvent, Data: req, SendAt: time.Now().Unix()}
	//通知消息服务有消息需要发送
	err = rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
	if err != nil {
		return
	}

	response.Success(c, "发送好友请求成功", nil)
}

// @Summary 群聊成员列表
// @Description 群聊成员列表
// @Param group_id query integer true "群聊ID"
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/group/member [get]
func getGroupMember(c *gin.Context) {
	// 从请求中获取群聊ID
	groupID := c.Query("group_id")
	if groupID == "" {
		response.Fail(c, "群聊ID不能为空", nil)
		return
	}

	gid, err := strconv.ParseUint(groupID, 10, 32)
	if err != nil {
		response.Fail(c, "群聊ID格式错误", nil)
		return
	}

	groupRelation, err := groupRelationClient.GetGroupUserIDs(context.Background(), &relationApi.GroupIDRequest{GroupId: uint32(gid)})
	if err != nil {
		response.Fail(c, "获取群聊成员失败", nil)
		return
	}

	resp, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: groupRelation.UserIds})
	if err != nil {
		c.Error(err)
		return
	}

	type requestListResponse struct {
		UserID   string `json:"user_id"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}

	var ids []string
	var data []*requestListResponse
	for _, v := range resp.Users {
		ids = append(ids, v.UserId)
		data = append(data, &requestListResponse{
			UserID:   v.UserId,
			Nickname: v.NickName,
			Avatar:   v.Avatar,
		})
	}

	response.Success(c, "获取群聊成员成功", data)
}

// groupRequestList 获取群聊申请列表
// @Summary 获取群聊申请列表
// @Description 获取用户的群聊申请列表 status 申请状态 (0=申请中, 1=已加入, 2=被拒绝, 3=被封禁)
// @Tags GroupRelation
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @Success		200 {object} model.Response{data=model.RequestListResponse}
// @Router /relation/group/request_list [get]
func groupRequestList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}
	// 检查用户是否存在
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		logger.Error("Failed to get user information", zap.Error(err))
		c.Error(err)
		return
	}

	reqList, err := groupRelationClient.GetUserGroupRequestList(context.Background(), &relationApi.GetUserGroupRequestListRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	var ids []string
	var data []*model.RequestListResponse
	for _, v := range reqList.GroupJoinRequestList {
		ids = append(ids, v.UserId)
		data = append(data, &model.RequestListResponse{
			UserID: v.UserId,
			Msg:    v.Msg,
			Status: uint32(v.Status),
			//RequestAt: v.CreatedAt,
		})
	}

	users, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: ids})
	if err != nil {
		c.Error(err)
		return
	}

	for _, v := range data {
		for _, u := range users.Users {
			if v.UserID == u.UserId {
				v.Nickname = u.NickName
				v.Avatar = u.Avatar
				break
			}
		}
	}

	response.Success(c, "获取群聊申请列表成功", data)
}

// @Summary 加入群聊
// @Description 加入群聊
// @Accept  json
// @Produce  json
// @param request body model.JoinGroupRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/group/join [post]
func joinGroup(c *gin.Context) {
	req := new(model.JoinGroupRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	uid, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	group, err := groupClient.GetGroupInfoByGid(context.Background(), &groupApi.GetGroupInfoRequest{Gid: req.GroupID})
	if err != nil {
		c.Error(err)
		return
	}

	if group.Status != groupApi.GroupStatus_GROUP_STATUS_NORMAL {
		response.Fail(c, "群聊状态不可用", nil)
		return
	}
	//判断是否在群聊中
	relation, err := groupRelationClient.GetGroupRelation(context.Background(), &relationApi.GetGroupRelationRequest{
		GroupId: req.GroupID,
		UserId:  uid,
	})
	if relation != nil {
		if relation.Status == relationApi.GroupRelationStatus_GroupStatusJoined {
			response.Fail(c, "您已经在群聊中", nil)
			return
		}
		if relation.Status == relationApi.GroupRelationStatus_GroupStatusReject {
			response.Fail(c, "拒绝加入群聊", nil)
			return
		}
		if relation.Status == relationApi.GroupRelationStatus_GroupStatusBlocked {
			response.Fail(c, "群聊状态不可用", nil)
			return
		}
	}

	//添加普通用户申请
	_, err = groupRelationClient.JoinGroup(context.Background(), &relationApi.JoinGroupRequest{UserId: uid, GroupId: req.GroupID, Identify: relationApi.GroupIdentity_IDENTITY_USER})
	if err != nil {
		c.Error(err)
		return
	}
	//查询所有管理员
	adminIds, err := groupRelationClient.GetGroupAdminIds(context.Background(), &relationApi.GroupIDRequest{
		GroupId: req.GroupID,
	})
	for _, id := range adminIds.UserIds {
		msg := msgconfig.WsMsg{Uid: id, Event: msgconfig.JoinGroupEvent, Data: map[string]interface{}{"group_id": req.GroupID, "user_id": uid}, SendAt: time.Now().Unix()}
		//通知消息服务有消息需要发送
		err = rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
		if err != nil {
			return
		}
	}

	response.Success(c, "发送加入群聊请求成功", nil)
}

// @Summary 管理加入群聊
// @Description 管理加入群聊
// @Accept  json
// @Produce  json
// @param request body model.ManageJoinGroupRequest true "Status (0: rejected, 1: joined)"
// @Success		200 {object} model.Response{}
// @Router /relation/group/manage_join_group [post]
func manageJoinGroup(c *gin.Context) {
	req := new(model.ManageJoinGroupRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	group, err := groupClient.GetGroupInfoByGid(context.Background(), &groupApi.GetGroupInfoRequest{Gid: req.GroupID})
	if err != nil {
		c.Error(err)
		return
	}

	if group.Status != groupApi.GroupStatus_GROUP_STATUS_NORMAL {
		response.Fail(c, "群聊状态不可用", nil)
		return
	}

	relation, err := groupRelationClient.GetGroupRelation(context.Background(), &relationApi.GetGroupRelationRequest{GroupId: req.GroupID, UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}
	if relation == nil {
		response.Fail(c, "没有加入群聊的申请", nil)
		return
	}
	if relation.Status != relationApi.GroupRelationStatus_GroupStatusJoined {
		response.Fail(c, "已经加入群聊", nil)
		return
	}
	id, err := dialogClient.GetDialogByGroupId(context.Background(), &relationApi.GetDialogByGroupIdRequest{GroupId: req.GroupID})
	if err != nil {
		c.Error(err)
		return
	}
	_, err = dialogClient.JoinDialog(context.Background(), &relationApi.JoinDialogRequest{DialogId: id.DialogId, UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}
	var status relationApi.GroupRelationStatus
	if req.Status == 1 {
		status = relationApi.GroupRelationStatus_GroupStatusJoined
	} else {
		status = relationApi.GroupRelationStatus_GroupStatusReject
	}
	// 执行同意加入群聊操作
	if _, err = groupRelationClient.ManageJoinGroup(context.Background(), &relationApi.ManageJoinGroupRequest{UserId: req.UserID, GroupId: req.GroupID, Status: status}); err != nil {
		c.Error(err)
		return
	}
	msg := msgconfig.WsMsg{Uid: req.UserID, Event: msgconfig.JoinGroupEvent, Data: map[string]interface{}{"group_id": req.GroupID, "status": status}, SendAt: time.Now().Unix()}
	//通知消息服务有消息需要发送
	err = rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
	if err != nil {
		return
	}
	response.Success(c, "管理群聊申请成功", nil)
}

// @Summary 将用户从群聊移除
// @Description 将用户从群聊移除
// @Accept  json
// @Produce  json
// @param request body model.RemoveUserFromGroupRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/group/remove [post]
func removeUserFromGroup(c *gin.Context) {
	req := new(model.RemoveUserFromGroupRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	if userID == req.UserID {
		response.Fail(c, "不能将自己从群聊中移除", nil)
		return
	}

	gr, err := groupRelationClient.GetGroupRelation(context.Background(), &relationApi.GetGroupRelationRequest{UserId: userID, GroupId: req.GroupID})
	if err != nil {
		c.Error(err)
		return
	}

	if gr.Identity != relationApi.GroupIdentity_IDENTITY_ADMIN {
		response.Fail(c, "没有权限操作", nil)
		return
	}

	if _, err = groupRelationClient.RemoveFromGroup(context.Background(), &relationApi.RemoveFromGroupRequest{UserId: userID, GroupId: req.GroupID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "移出群聊成功", nil)
}

// @Summary 退出群聊
// @Description 退出群聊
// @Accept  json
// @Produce  json
// @param request body model.QuitGroupRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/group/quit [post]
func quitGroup(c *gin.Context) {
	req := new(model.QuitGroupRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	//查询用户是否在群聊中
	if _, err = groupRelationClient.GetGroupRelation(context.Background(), &relationApi.GetGroupRelationRequest{UserId: userID, GroupId: req.GroupID}); err != nil {
		c.Error(err)
		return
	}
	//删除用户对话
	if _, err = dialogClient.DeleteDialogUserByDialogIDAndUserID(context.Background(), &relationApi.DeleteDialogUserByDialogIDAndUserIDRequest{
		DialogId: req.GroupID,
		UserId:   userID,
	}); err != nil {
		c.Error(err)
		return
	}
	//退出群聊
	if _, err = groupRelationClient.LeaveGroup(context.Background(), &relationApi.LeaveGroupRequest{UserId: userID, GroupId: req.GroupID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "退出群聊成功", nil)
}

// @Summary 交换用户端到端公钥
// @Description 交换用户端到端公钥
// @Accept json
// @Produce json
// @param request body model.SwitchUserE2EPublicKeyRequest true "request"
// @Security BearerToken
// @Success 200 {object} model.Response{}
// @Router /relation/user/switch/e2e/key [post]
func switchUserE2EPublicKey(c *gin.Context) {
	req := new(model.SwitchUserE2EPublicKeyRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	// 获取用户ID，可以从请求中的token中解析出来，前提是你的登录接口已经设置了正确的token
	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	_, err = userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserId})
	if err != nil {
		response.Fail(c, "用户不存在", nil)
		return
	}
	reqm := model.SwitchUserE2EPublicKeyRequest{
		UserId:    thisId,
		PublicKey: req.PublicKey,
	}
	msg := msgconfig.WsMsg{Uid: req.UserId, Event: msgconfig.PushE2EPublicKeyEvent, Data: reqm, SendAt: time.Now().Unix()}

	//通知消息服务有消息需要发送
	err = rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
	if err != nil {
		return
	}

	response.Success(c, "交换用户公钥成功", nil)
}
