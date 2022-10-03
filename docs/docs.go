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
            "name": "@aasumitro",
            "url": "https://aasumitro.id/",
            "email": "hello@aasumitro.id"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/aasumitro/karlota/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/login": {
            "post": {
                "description": "Generate Access Token (JWT).",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountHandler"
                ],
                "summary": "Logged User In",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "CREATED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "422": {
                        "description": "UNPROCESSABLE_ENTITY_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ValidationErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        },
        "/v1/profile": {
            "get": {
                "description": "Get User Data in Detail.",
                "tags": [
                    "AccountHandler"
                ],
                "summary": "User Profile",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "401": {
                        "description": "UNAUTHORIZED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "description": "Generate new User Account.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountHandler"
                ],
                "summary": "Register new User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "full name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "CREATED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "422": {
                        "description": "UNPROCESSABLE_ENTITY_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ValidationErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        },
        "/v1/update/fcm": {
            "post": {
                "description": "Store Firebase Cloud Messaging Token.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountHandler"
                ],
                "summary": "Update FCM TOKEN",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "firebase cloud messaging token",
                        "name": "fcm_token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "CREATED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "422": {
                        "description": "UNPROCESSABLE_ENTITY_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ValidationErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        },
        "/v1/update/password": {
            "post": {
                "description": "Generate New Password.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountHandler"
                ],
                "summary": "Update Password",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "new password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "CREATED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "422": {
                        "description": "UNPROCESSABLE_ENTITY_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ValidationErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Get User List.",
                "tags": [
                    "AccountHandler"
                ],
                "summary": "User List",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.SuccessRespond"
                        }
                    },
                    "400": {
                        "description": "BAD_REQUEST_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "401": {
                        "description": "UNAUTHORIZED_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR_RESPOND",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorRespond"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.ErrorRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "utils.SuccessRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "utils.ValidationErrorRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
