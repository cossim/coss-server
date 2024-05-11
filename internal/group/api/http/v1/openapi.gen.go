// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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
	// 获取所有群聊信息
	// (GET /api/v1/group)
	GetAllGroup(c *gin.Context)
	// 创建群聊
	// (POST /api/v1/group)
	CreateGroup(c *gin.Context)
	// 搜索群聊
	// (GET /api/v1/group/search)
	SearchGroup(c *gin.Context, params SearchGroupParams)
	// 删除群聊
	// (DELETE /api/v1/group/{id})
	DeleteGroup(c *gin.Context, id uint32)
	// 获取群聊信息
	// (GET /api/v1/group/{id})
	GetGroup(c *gin.Context, id uint32)
	// 更新群聊信息
	// (PUT /api/v1/group/{id})
	UpdateGroup(c *gin.Context, id uint32)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAllGroup operation middleware
func (siw *ServerInterfaceWrapper) GetAllGroup(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAllGroup(c)
}

// CreateGroup operation middleware
func (siw *ServerInterfaceWrapper) CreateGroup(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateGroup(c)
}

// SearchGroup operation middleware
func (siw *ServerInterfaceWrapper) SearchGroup(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchGroupParams

	// ------------- Required query parameter "keyword" -------------

	if paramValue := c.Query("keyword"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument keyword is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "keyword", c.Request.URL.Query(), &params.Keyword)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter keyword: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", c.Request.URL.Query(), &params.PageSize)
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

	siw.Handler.SearchGroup(c, params)
}

// DeleteGroup operation middleware
func (siw *ServerInterfaceWrapper) DeleteGroup(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteGroup(c, id)
}

// GetGroup operation middleware
func (siw *ServerInterfaceWrapper) GetGroup(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGroup(c, id)
}

// UpdateGroup operation middleware
func (siw *ServerInterfaceWrapper) UpdateGroup(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateGroup(c, id)
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

	router.GET(options.BaseURL+"/api/v1/group", wrapper.GetAllGroup)
	router.POST(options.BaseURL+"/api/v1/group", wrapper.CreateGroup)
	router.GET(options.BaseURL+"/api/v1/group/search", wrapper.SearchGroup)
	router.DELETE(options.BaseURL+"/api/v1/group/:id", wrapper.DeleteGroup)
	router.GET(options.BaseURL+"/api/v1/group/:id", wrapper.GetGroup)
	router.PUT(options.BaseURL+"/api/v1/group/:id", wrapper.UpdateGroup)
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
