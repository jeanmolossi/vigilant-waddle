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
        "termsOfService": "github.com/jeanmolossi/vigilant-waddle/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/me": {
            "get": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Get current student",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Get current student",
                "operationId": "get-me",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "fields to return from the student",
                        "name": "fields",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpNewStudent"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPBadRequestError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPError"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPError"
                        }
                    },
                    "500": {
                        "description": "An error occurred",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPError"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "A simple health check.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Ping the server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    }
                }
            }
        },
        "/student": {
            "post": {
                "description": "Register a student",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Register a student",
                "operationId": "register-student",
                "parameters": [
                    {
                        "description": "Student",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/adapter.RegisterStudent"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpNewStudent"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPBadRequestError"
                        }
                    },
                    "409": {
                        "description": "User with that email already exists",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPError"
                        }
                    },
                    "500": {
                        "description": "An error occurred",
                        "schema": {
                            "$ref": "#/definitions/http_error.HTTPError"
                        }
                    }
                }
            }
        },
        "/students": {
            "get": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Get all students",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Get all students",
                "operationId": "get-students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {}
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "adapter.RegisterStudent": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6,
                    "example": "123456"
                }
            }
        },
        "handler.HttpNewStudent": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handler.HttpStudent"
                }
            }
        },
        "handler.HttpStudent": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "format": "email",
                    "example": "john@doe.com"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "http_error.HTTPBadRequestError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validator.FieldError"
                    }
                }
            }
        },
        "http_error.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "httputil.PingInternalServerErr": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "unexpected error"
                }
            }
        },
        "httputil.PingOk": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "validator.FieldError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "field_name"
                },
                "message": {
                    "type": "string",
                    "example": "field_name is required"
                }
            }
        }
    },
    "securityDefinitions": {
        "access_token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
