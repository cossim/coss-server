// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplategroup = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/group/create": {
            "post": {
                "description": "创建群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建群聊",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/group/delete": {
            "post": {
                "description": "删除群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除群聊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "群聊ID",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/group/getBatch": {
            "get": {
                "description": "批量获取群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "批量获取群聊信息",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "群聊ID列表",
                        "name": "groupIds",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/group/info": {
            "get": {
                "description": "获取群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取群聊信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "群聊ID",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/group/update/{gid}": {
            "post": {
                "description": "更新群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新群聊信息",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.createGroupRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "max_members_limit": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "http.updateGroupRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "max_members_limit": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfogroup holds exported Swagger Info so clients can modify it
var SwaggerInfogroup = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "coss-user服务",
	Description:      "",
	InfoInstanceName: "group",
	SwaggerTemplate:  docTemplategroup,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfogroup.InstanceName(), SwaggerInfogroup)
}
