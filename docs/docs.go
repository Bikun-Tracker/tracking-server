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
        "/bus/": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bus"
                ],
                "summary": "Create new bus entry",
                "parameters": [
                    {
                        "description": "CreateBus",
                        "name": "CreateBusDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusResponse"
                        }
                    }
                }
            }
        },
        "/bus/login": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bus"
                ],
                "summary": "Driver login",
                "parameters": [
                    {
                        "description": "DriverLoginDto",
                        "name": "DriverLoginDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DriverLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DriverLoginResponse"
                        }
                    }
                }
            }
        },
        "/bus/{id}": {
            "put": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bus"
                ],
                "summary": "Edit Bus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bus ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "EditBusDto",
                        "name": "EditBusDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EditBusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EditBusResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bus"
                ],
                "summary": "Delete bus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bus ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "Check system status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Status"
                            }
                        }
                    }
                }
            }
        },
        "/news/": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "Get all news",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetAllNewsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "Create new news",
                "parameters": [
                    {
                        "description": "CreateNews",
                        "name": "CreateNewsDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewsDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateBusDto": {
            "type": "object",
            "required": [
                "number",
                "password",
                "plate",
                "route",
                "username"
            ],
            "properties": {
                "number": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "plate": {
                    "type": "string"
                },
                "route": {
                    "type": "string",
                    "enum": [
                        "RED",
                        "BLUE"
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CreateBusResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "plate": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CreateNewsDto": {
            "type": "object",
            "required": [
                "detail",
                "title"
            ],
            "properties": {
                "detail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.CreateNewsResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.DriverLoginDto": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.DriverLoginResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "plate": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.EditBusDto": {
            "type": "object",
            "properties": {
                "isActive": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "plate": {
                    "type": "string"
                },
                "route": {
                    "type": "string",
                    "enum": [
                        "RED",
                        "BLUE"
                    ]
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "EMPTY",
                        "MODERATE",
                        "FULL"
                    ]
                }
            }
        },
        "dto.EditBusResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "plate": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.GetAllNewsResponse": {
            "type": "object",
            "properties": {
                "news": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.News"
                    }
                }
            }
        },
        "dto.News": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.Status": {
            "type": "object",
            "properties": {
                "data": {},
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Bikun Tracking API",
	Description:      "API definition for bikun tracking specification",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
