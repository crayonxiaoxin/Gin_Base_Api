// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/login": {
            "post": {
                "description": "登入",
                "tags": [
                    "login"
                ],
                "summary": "登入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "user"
                        }
                    }
                }
            }
        },
        "/media": {
            "get": {
                "description": "获取所有媒体",
                "tags": [
                    "media"
                ],
                "summary": "获取所有媒体",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            }
        },
        "/media/{id}": {
            "get": {
                "description": "通过id获取文件",
                "tags": [
                    "media"
                ],
                "summary": "通过id获取文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The key for staticblock",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除文件",
                "tags": [
                    "media"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "文件ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册",
                "tags": [
                    "register"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "user"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "上传",
                "tags": [
                    "media"
                ],
                "summary": "上传",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "获取所有用户",
                "tags": [
                    "user"
                ],
                "summary": "获取所有用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            },
            "post": {
                "description": "添加用户",
                "tags": [
                    "user"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "通过id获取用户",
                "tags": [
                    "user"
                ],
                "summary": "通过id获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The key for staticblock",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            },
            "put": {
                "description": "更新用户",
                "tags": [
                    "user"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The uid you want to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户",
                "tags": [
                    "user"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登入后返回的token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The uid you want to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.Result": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "rc": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Test API",
	Description:      "Gin API 基础工程",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
