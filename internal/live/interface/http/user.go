package http

import (
	"github.com/cossim/coss-server/internal/live/api/dto"
	pkghttp "github.com/cossim/coss-server/pkg/http"
	"github.com/cossim/coss-server/pkg/http/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserCreate
// @Summary 创建用户通话
// @Description 创建用户通话
// @Tags LiveUser
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @param request body dto.UserCallRequest true "request"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response{data=dto.UserCallResponse}
// @Router /live/user/create [post]
func (h *Handler) UserCreate(c *gin.Context) {
	req := new(dto.UserCallRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数验证失败", nil)
		return
	}

	userID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.CreateUserCall(c, userID, req)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "创建通话成功", resp)
}

// UserJoin
// @Summary 加入通话
// @Description 加入通话
// @Tags LiveUser
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @param request body dto.UserJoinRequest true "request"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response{data=dto.UserJoinResponse}
// @Router /live/user/join [post]
func (h *Handler) UserJoin(c *gin.Context) {
	req := new(dto.UserJoinRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetFail(c, "参数验证失败", nil)
		return
	}

	userID, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		h.logger.Error("token解析失败", zap.Error(err))
		response.Fail(c, "token解析失败", nil)
		return
	}

	driverId, err := pkghttp.ParseTokenReDriverId(c)
	if err != nil {
		c.Error(err)
		return
	}

	resp, err := h.svc.UserJoinRoom(c, userID, driverId)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "加入通话成功", resp)
}

// UserShow
// @Summary 获取通话房间信息
// @Description 获取通话房间信息
// @Tags LiveUser
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @Produce  json
// @Success		200 {object} dto.Response{data=dto.UserShowResponse} "participant=通话参与者"
// @Router /live/user/show [get]
func (h *Handler) UserShow(c *gin.Context) {
	uid, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, "token解析失败", nil)
		return
	}

	resp, err := h.svc.GetUserRoom(c, uid)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "获取通话信息成功", resp)
}

// UserReject
// @Summary 拒绝通话
// @Description 拒绝通话
// @Tags LiveUser
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @Accept  json
// @Produce  json
// @Success		200 {object} dto.Response{}
// @Router /live/user/reject [post]
func (h *Handler) UserReject(c *gin.Context) {
	//req := new(dto.UserRejectRequest)
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	response.SetFail(c, "参数验证失败", nil)
	//	return
	//}

	uid, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, "token解析失败", nil)
		return
	}

	driverId, err := pkghttp.ParseTokenReDriverId(c)
	if err != nil {
		c.Error(err)
		return
	}

	resp, err := h.svc.UserRejectRoom(c, uid, driverId)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "拒绝通话成功", resp)
}

// UserLeave
// @Summary 结束通话
// @Description 结束通话
// @Tags LiveUser
// @Security Bearer
// @Param Authorization header string true "Bearer JWT"
// @Accept  json
// @Produce  json
// @Success		200 {object} dto.Response{}
// @Router /live/user/leave [post]
func (h *Handler) UserLeave(c *gin.Context) {
	//req := new(dto.UserLeaveRequest)
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	response.SetFail(c, "参数验证失败", nil)
	//	return
	//}

	uid, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.SetFail(c, "token解析失败", nil)
		return
	}

	driverId, err := pkghttp.ParseTokenReDriverId(c)
	if err != nil {
		c.Error(err)
		return
	}

	resp, err := h.svc.UserLeaveRoom(c, uid, driverId)
	if err != nil {
		c.Error(err)
		return
	}

	response.SetSuccess(c, "结束通话成功", resp)
}

func (h *Handler) getJoinToken(c *gin.Context) {
	//userID := c.Query("user_id")
	//if userID == "" {
	//	response.SetFail(c, code.InvalidParameter.Message(), nil)
	//	return
	//}
	//
	//resp, err := h.svc.GetAdminJoinToken(c, userID, "")
	//if err != nil {
	//	c.Error(err)
	//	return
	//}
	//
	//response.SetSuccess(c, "获取token成功", resp)
}
