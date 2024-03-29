package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

const (
	InternalServerErrorCode = http.StatusInternalServerError
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	ctx.JSON(httpStatus, BaseResponse{
		Code: code,
		Msg:  extractErrorMessage(msg),
		Data: data,
	})
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

func Fail(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, 400, msg, data)
}

func InternalServerError(ctx *gin.Context) {
	Response(ctx, InternalServerErrorCode, InternalServerErrorCode, http.StatusText(InternalServerErrorCode), nil)
}

func NetworkErrorRetry(ctx *gin.Context) {
	Response(ctx, 400, 400, "Network error. Please retry.", nil)
}

func GRPCError(ctx *gin.Context, err error) {
	InternalServerError(ctx)
}

func extractErrorMessage(input string) string {
	re := regexp.MustCompile(`rpc error: code = \w+ desc = (.+)`)
	match := re.FindStringSubmatch(input)

	if len(match) < 2 {
		return input
	}

	return match[1]
}

func SetResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.Set("response", BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func SetSuccess(ctx *gin.Context, msg string, data interface{}) {
	SetResponse(ctx, 200, msg, data)
}

func SetFail(ctx *gin.Context, msg string, data interface{}) {
	SetResponse(ctx, 400, msg, data)
}

func SetInternalServerError(ctx *gin.Context) {
	SetResponse(ctx, InternalServerErrorCode, http.StatusText(InternalServerErrorCode), nil)
}
