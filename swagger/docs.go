// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
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
        "/authentication/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.AuthenticationLoginRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.AuthenticationLoginApiResponse"
                        }
                    }
                }
            }
        },
        "/authentication/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.AuthenticationRegisterRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.AuthenticationRegisterApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.AuthenticationLoginRequestData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.AuthenticationLoginResponseData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dtos.AuthenticationRegisterRequestData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.AuthenticationRegisterResponseData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "swagger.AuthenticationLoginApiResponse": {
            "type": "object",
            "properties": {
                "additionalInfo": {
                    "type": "object",
                    "additionalProperties": true
                },
                "data": {
                    "$ref": "#/definitions/dtos.AuthenticationLoginResponseData"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "swagger.AuthenticationRegisterApiResponse": {
            "type": "object",
            "properties": {
                "additionalInfo": {
                    "type": "object",
                    "additionalProperties": true
                },
                "data": {
                    "$ref": "#/definitions/dtos.AuthenticationRegisterResponseData"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Api Template",
	Description:      "This is a template for a RESTful API using Fiber.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}