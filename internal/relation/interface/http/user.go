package http

import (
	"github.com/cossim/coss-server/internal/relation/api/http/model"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/http"
	pkghttp "github.com/cossim/coss-server/pkg/http"
	"github.com/cossim/coss-server/pkg/http/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 设置好友阅后即焚消息销毁时间
// @Summary 设置好友阅后即焚消息销毁时间
// @Description 设置好友阅后即焚消息销毁时间
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.SetUserOpenBurnAfterReadingTimeOutRequest true "request"
// @Success 200 {object} model.Response{}
// @Router /relation/user/burn/timeout/set [post]
func (h *Handler) setUserOpenBurnAfterReadingTimeOut(c *gin.Context) {
	req := new(model.SetUserOpenBurnAfterReadingTimeOutRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}
	thisID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}

	err = h.svc.SetUserOpenBurnAfterReadingTimeOut(c, thisID, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "设置成功", nil)
}

// 修改好友备注
// @Summary 修改好友备注
// @Description 修改用户备注
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.SetUserFriendRemarkRequest true "request"
// @Success 200 {object} model.Response{}
// @Router /relation/user/remark/set [post]
func (h *Handler) setUserFriendRemark(c *gin.Context) {
	req := new(model.SetUserFriendRemarkRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}
	thisID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}

	_, err = h.svc.SetUserFriendRemark(c, thisID, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "修改成功", nil)
}

// @Summary 是否打开用户阅后即焚消息(action: 0:关闭, 1:打开)
// @Description 是否打开用户阅后即焚消息(action: 0:关闭, 1:打开)
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.OpenUserBurnAfterReadingRequest true "request"
// @Success 200 {object} model.Response{}
// @Router /relation/user/burn/open [post]
func (h *Handler) openUserBurnAfterReading(c *gin.Context) {
	req := new(model.OpenUserBurnAfterReadingRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}

	if !model.IsValidOpenBurnAfterReadingType(req.Action) {
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	_, err = h.svc.SetUserBurnAfterReading(c, thisId, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "操作成功", nil)
}

// @Summary 设置私聊静默通知
// @Description 设置私聊静默通知
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.SetUserSilentNotificationRequest true "request"
// @Success 200 {object} model.Response{}
// @Router /relation/user/silent [post]
func (h *Handler) setUserSilentNotification(c *gin.Context) {
	req := new(model.SetUserSilentNotificationRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}

	if !model.IsValidSilentNotificationType(req.IsSilent) {
		response.SetFail(c, "设置消息标注状态失败", nil)
		return
	}

	_, err = h.svc.UserSilentNotification(c, thisId, req.UserId, req.IsSilent)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "设置成功", gin.H{"user_id": req.UserId})
}

// @Summary 交换用户端到端公钥
// @Description 交换用户端到端公钥
// @Tags UserRelation
// @Accept json
// @Produce json
// @param request body model.SwitchUserE2EPublicKeyRequest true "request"
// @Security BearerToken
// @Success 200 {object} model.Response{}
// @Router /relation/user/switch/e2e/key [post]
func (h *Handler) switchUserE2EPublicKey(c *gin.Context) {
	req := new(model.SwitchUserE2EPublicKeyRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	// 获取用户ID，可以从请求中的token中解析出来，前提是你的登录接口已经设置了正确的token
	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}

	_, err = h.svc.SwitchUserE2EPublicKey(c, thisId, req.UserId, req.PublicKey)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "交换用户公钥成功", nil)
}

// @Summary 黑名单
// @Description 黑名单
// @Tags UserRelation
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/user/blacklist [get]
func (h *Handler) blackList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.BlackList(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取黑名单列表成功", resp)
}

// @Summary 好友列表
// @Description 好友列表
// @Tags UserRelation
// @Produce  json
// @Success		200 {object} model.Response{}
// @Router /relation/user/friend_list [get]
func (h *Handler) friendList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.FriendList(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取好友列表成功", resp)
}

// @Summary 群聊列表
// @Description 群聊列表
// @Tags GroupRelation
// @Produce  json
// @Success		200 {object} model.Response{data=[]usersorter.CustomGroupData} "status 0:正常状态；1:被封禁状态；2:被删除状态"
// @Router /relation/group/list [get]
func (h *Handler) getUserGroupList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.GetUserGroupList(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取用户群聊列表成功", resp)
}

// @Summary 好友申请列表
// @Description 好友申请列表
// @Tags UserRelation
// @Produce  json
// @Success		200 {object} model.Response{data=[]model.UserRequestListResponse} "UserStatus 申请状态 (0=申请中, 1=待通过, 2=已添加, 3=被拒绝, 4=已删除, 5=已拒绝)"
// @Router /relation/user/request_list [get]
func (h *Handler) userRequestList(c *gin.Context) {
	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.UserRequestList(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取好友申请列表成功", resp)
}

// @Summary 删除黑名单
// @Description 删除黑名单
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.DeleteBlacklistRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/delete_blacklist [post]
func (h *Handler) deleteBlacklist(c *gin.Context) {
	req := new(model.DeleteBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	_, err = h.svc.DeleteBlacklist(c, userID, req.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "删除黑名单成功", nil)
}

// @Summary 添加黑名单
// @Description 添加黑名单
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.AddBlacklistRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/add_blacklist [post]
func (h *Handler) addBlacklist(c *gin.Context) {
	req := new(model.AddBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	_, err = h.svc.AddBlacklist(c, userID, req.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "添加到黑名单成功", nil)
}

// @Summary 删除好友
// @Description 删除好友
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.DeleteFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/delete_friend [post]
func (h *Handler) deleteFriend(c *gin.Context) {
	req := new(model.DeleteFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	if err = h.svc.DeleteFriend(c, userID, req.UserID); err != nil {
		response.SetFail(c, "删除好友失败", nil)
		return
	}

	response.SetSuccess(c, "删除好友成功", nil)
}

// @Summary 管理好友请求
// @Description 管理好友请求  action (0=拒绝, 1=同意)
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.ManageFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/manage_friend [post]
func (h *Handler) manageFriend(c *gin.Context) {
	req := new(model.ManageFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	if err := req.Validator(); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	responseData, err := h.svc.ManageFriend(c, userID, req.RequestID, int32(req.Action), req.E2EPublicKey)
	if err != nil {
		response.SetFail(c, code.Cause(err).Message(), nil)
		return
	}

	response.SetSuccess(c, "管理好友申请成功", responseData)
}

// @Summary 发送好友请求
// @Description 发送好友请求
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.SendFriendRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/add_friend [post]
func (h *Handler) addFriend(c *gin.Context) {
	req := new(model.SendFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}
	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, err.Error(), nil)
		return
	}
	resp, err := h.svc.SendFriendRequest(c, thisId, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "发送好友请求成功", resp)
}

// @Summary 删除好友申请记录
// @Description 删除好友申请记录
// @Tags UserRelation
// @Accept  json
// @Produce  json
// @param request body model.DeleteRecordRequest true "request"
// @Success		200 {object} model.Response{}
// @Router /relation/user/delete_friend_record [post]
func (h *Handler) deleteUserFriendRecord(c *gin.Context) {
	req := new(model.DeleteRecordRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("参数验证失败", zap.Error(err))
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := http.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.SetFail(c, "token解析失败", nil)
		return
	}

	if err = h.svc.DeleteUserFriendRecord(c, userID, req.ID); err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "删除好友申请记录成功", nil)
}
