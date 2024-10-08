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
        "contact": {
            "name": "Api service support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello_world": {
            "get": {
                "security": [
                    {
                        "User": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "Get hello world test info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.HelloWorld"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.HelloWorld": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "response.errorResp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Internal apperror code",
                    "type": "integer"
                },
                "error": {
                    "description": "Error message to display",
                    "type": "string"
                },
                "request_id": {
                    "description": "id to determinate what exacly was wrong by searching in logs.",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "UserToken": {
            "type": "apiKey",
            "name": "X-Token-Key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Basic API",
	Description:      "Basic api swagger description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
