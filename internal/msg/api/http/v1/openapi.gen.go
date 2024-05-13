// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 获取指定对话中延迟的消息
	// (POST /api/v1/msg/dialog/after)
	GetAfterMsgs(c *gin.Context)
	// 获取群组标记消息列表
	// (GET /api/v1/msg/dialog/group/{dialog_id}/label)
	GetGroupLabelMsgList(c *gin.Context, dialogId int)
	// 获取用户对话列表
	// (GET /api/v1/msg/dialog/list)
	GetUserDialogList(c *gin.Context, params GetUserDialogListParams)
	// 获取用户标记消息列表
	// (GET /api/v1/msg/dialog/user/{dialog_id}/label)
	GetUserLabelMsgList(c *gin.Context, dialogId int)
	// 获取群组消息列表
	// (GET /api/v1/msg/group/list)
	GetGroupMsgList(c *gin.Context, params GetGroupMsgListParams)
	// 设置群组消息已读
	// (PUT /api/v1/msg/group/read)
	GroupMessageRead(c *gin.Context)
	// 发送群组消息
	// (POST /api/v1/msg/group/send)
	SendGroupMsg(c *gin.Context)
	// 撤回群组消息
	// (DELETE /api/v1/msg/group/{id})
	RecallGroupMsg(c *gin.Context, id int)
	// 编辑群组消息
	// (PUT /api/v1/msg/group/{id})
	EditGroupMsg(c *gin.Context, id int)
	// 标记群组消息
	// (POST /api/v1/msg/group/{id}/label)
	LabelGroupMsg(c *gin.Context, id int)
	// 获取群组消息阅读者
	// (GET /api/v1/msg/group/{id}/read)
	GetGroupMessageReaders(c *gin.Context, id int, params GetGroupMessageReadersParams)
	// 获取私信列表
	// (GET /api/v1/msg/user/list)
	GetUserMsgList(c *gin.Context, params GetUserMsgListParams)
	// 设置用户消息已读
	// (PUT /api/v1/msg/user/read)
	ReadUserMsgs(c *gin.Context)
	// 发送私信
	// (POST /api/v1/msg/user/send)
	SendUserMsg(c *gin.Context)
	// 撤回用户消息
	// (DELETE /api/v1/msg/user/{id})
	RecallUserMsg(c *gin.Context, id int)
	// 编辑用户消息
	// (PUT /api/v1/msg/user/{id})
	EditUserMsg(c *gin.Context, id int)
	// 标记用户消息
	// (POST /api/v1/msg/user/{id}/label)
	LabelUserMsg(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAfterMsgs operation middleware
func (siw *ServerInterfaceWrapper) GetAfterMsgs(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAfterMsgs(c)
}

// GetGroupLabelMsgList operation middleware
func (siw *ServerInterfaceWrapper) GetGroupLabelMsgList(c *gin.Context) {

	var err error

	// ------------- Path parameter "dialog_id" -------------
	var dialogId int

	err = runtime.BindStyledParameter("simple", false, "dialog_id", c.Param("dialog_id"), &dialogId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter dialog_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGroupLabelMsgList(c, dialogId)
}

// GetUserDialogList operation middleware
func (siw *ServerInterfaceWrapper) GetUserDialogList(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserDialogListParams

	// ------------- Required query parameter "page_num" -------------

	if paramValue := c.Query("page_num"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_num is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_num", c.Request.URL.Query(), &params.PageNum)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_num: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "page_size" -------------

	if paramValue := c.Query("page_size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_size", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_size: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserDialogList(c, params)
}

// GetUserLabelMsgList operation middleware
func (siw *ServerInterfaceWrapper) GetUserLabelMsgList(c *gin.Context) {

	var err error

	// ------------- Path parameter "dialog_id" -------------
	var dialogId int

	err = runtime.BindStyledParameter("simple", false, "dialog_id", c.Param("dialog_id"), &dialogId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter dialog_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserLabelMsgList(c, dialogId)
}

// GetGroupMsgList operation middleware
func (siw *ServerInterfaceWrapper) GetGroupMsgList(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGroupMsgListParams

	// ------------- Required query parameter "dialog_id" -------------

	if paramValue := c.Query("dialog_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument dialog_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "dialog_id", c.Request.URL.Query(), &params.DialogId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter dialog_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "msg_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "msg_id", c.Request.URL.Query(), &params.MsgId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter msg_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "user_id", c.Request.URL.Query(), &params.UserId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", c.Request.URL.Query(), &params.Type)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter type: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "content" -------------

	err = runtime.BindQueryParameter("form", true, false, "content", c.Request.URL.Query(), &params.Content)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter content: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "page_num" -------------

	if paramValue := c.Query("page_num"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_num is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_num", c.Request.URL.Query(), &params.PageNum)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_num: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "page_size" -------------

	if paramValue := c.Query("page_size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_size", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_size: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGroupMsgList(c, params)
}

// GroupMessageRead operation middleware
func (siw *ServerInterfaceWrapper) GroupMessageRead(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GroupMessageRead(c)
}

// SendGroupMsg operation middleware
func (siw *ServerInterfaceWrapper) SendGroupMsg(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SendGroupMsg(c)
}

// RecallGroupMsg operation middleware
func (siw *ServerInterfaceWrapper) RecallGroupMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RecallGroupMsg(c, id)
}

// EditGroupMsg operation middleware
func (siw *ServerInterfaceWrapper) EditGroupMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.EditGroupMsg(c, id)
}

// LabelGroupMsg operation middleware
func (siw *ServerInterfaceWrapper) LabelGroupMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.LabelGroupMsg(c, id)
}

// GetGroupMessageReaders operation middleware
func (siw *ServerInterfaceWrapper) GetGroupMessageReaders(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGroupMessageReadersParams

	// ------------- Required query parameter "dialog_id" -------------

	if paramValue := c.Query("dialog_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument dialog_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "dialog_id", c.Request.URL.Query(), &params.DialogId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter dialog_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "group_id" -------------

	if paramValue := c.Query("group_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument group_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "group_id", c.Request.URL.Query(), &params.GroupId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter group_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGroupMessageReaders(c, id, params)
}

// GetUserMsgList operation middleware
func (siw *ServerInterfaceWrapper) GetUserMsgList(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserMsgListParams

	// ------------- Required query parameter "dialog_id" -------------

	if paramValue := c.Query("dialog_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument dialog_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "dialog_id", c.Request.URL.Query(), &params.DialogId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter dialog_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", c.Request.URL.Query(), &params.Type)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter type: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "user_id", c.Request.URL.Query(), &params.UserId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "content" -------------

	err = runtime.BindQueryParameter("form", true, false, "content", c.Request.URL.Query(), &params.Content)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter content: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "msg_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "msg_id", c.Request.URL.Query(), &params.MsgId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter msg_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "page_num" -------------

	if paramValue := c.Query("page_num"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_num is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_num", c.Request.URL.Query(), &params.PageNum)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_num: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "page_size" -------------

	if paramValue := c.Query("page_size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page_size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_size", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page_size: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_at" -------------

	err = runtime.BindQueryParameter("form", true, false, "start_at", c.Request.URL.Query(), &params.StartAt)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter start_at: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "end_at" -------------

	if paramValue := c.Query("end_at"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument end_at is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "end_at", c.Request.URL.Query(), &params.EndAt)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter end_at: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserMsgList(c, params)
}

// ReadUserMsgs operation middleware
func (siw *ServerInterfaceWrapper) ReadUserMsgs(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ReadUserMsgs(c)
}

// SendUserMsg operation middleware
func (siw *ServerInterfaceWrapper) SendUserMsg(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SendUserMsg(c)
}

// RecallUserMsg operation middleware
func (siw *ServerInterfaceWrapper) RecallUserMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RecallUserMsg(c, id)
}

// EditUserMsg operation middleware
func (siw *ServerInterfaceWrapper) EditUserMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.EditUserMsg(c, id)
}

// LabelUserMsg operation middleware
func (siw *ServerInterfaceWrapper) LabelUserMsg(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.LabelUserMsg(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/api/v1/msg/dialog/after", wrapper.GetAfterMsgs)
	router.GET(options.BaseURL+"/api/v1/msg/dialog/group/:dialog_id/label", wrapper.GetGroupLabelMsgList)
	router.GET(options.BaseURL+"/api/v1/msg/dialog/list", wrapper.GetUserDialogList)
	router.GET(options.BaseURL+"/api/v1/msg/dialog/user/:dialog_id/label", wrapper.GetUserLabelMsgList)
	router.GET(options.BaseURL+"/api/v1/msg/group/list", wrapper.GetGroupMsgList)
	router.PUT(options.BaseURL+"/api/v1/msg/group/read", wrapper.GroupMessageRead)
	router.POST(options.BaseURL+"/api/v1/msg/group/send", wrapper.SendGroupMsg)
	router.DELETE(options.BaseURL+"/api/v1/msg/group/:id", wrapper.RecallGroupMsg)
	router.PUT(options.BaseURL+"/api/v1/msg/group/:id", wrapper.EditGroupMsg)
	router.POST(options.BaseURL+"/api/v1/msg/group/:id/label", wrapper.LabelGroupMsg)
	router.GET(options.BaseURL+"/api/v1/msg/group/:id/read", wrapper.GetGroupMessageReaders)
	router.GET(options.BaseURL+"/api/v1/msg/user/list", wrapper.GetUserMsgList)
	router.PUT(options.BaseURL+"/api/v1/msg/user/read", wrapper.ReadUserMsgs)
	router.POST(options.BaseURL+"/api/v1/msg/user/send", wrapper.SendUserMsg)
	router.DELETE(options.BaseURL+"/api/v1/msg/user/:id", wrapper.RecallUserMsg)
	router.PUT(options.BaseURL+"/api/v1/msg/user/:id", wrapper.EditUserMsg)
	router.POST(options.BaseURL+"/api/v1/msg/user/:id/label", wrapper.LabelUserMsg)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}