// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/commands/create": {
            "post": {
                "summary": "Create and execute new command",
                "tags": [
                    "Commands"
                ],
                "operationId": "CreateCommand",
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Creation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful post-request",
                        "schema": {
                            "$ref": "#/definitions/CreationResponse"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful post-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/commands/run/{id}": {
            "post": {
                "summary": "Execute command with given id",
                "tags": [
                    "Commands"
                ],
                "operationId": "RunCommand",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "description": "Command's identifier",
                        "example": "1"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful post-request",
                        "schema": {
                            "$ref": "#/definitions/RunResponse"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful post-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/commands/list": {
            "get": {
                "summary": "Get list of available commands",
                "tags": [
                    "Commands"
                ],
                "operationId": "ListCommands",
                "responses": {
                    "200": {
                        "description": "Successful get-request",
                        "schema": {
                            "$ref": "#/definitions/Commands"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful get-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/commands/{id}": {
            "get": {
                "summary": "Get command with given id",
                "tags": [
                    "Commands"
                ],
                "operationId": "GetCommand",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "description": "Command's identifier",
                        "example": "1"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful get-request",
                        "schema": {
                            "$ref": "#/definitions/Command"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful get-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/commands/delete": {
            "delete": {
                "summary": "Delete all rows in table Commands",
                "tags": [
                    "DataBase"
                ],
                "operationId": "DeleteAllRows",
                "responses": {
                    "200": {
                        "description": "Successful delete-request",
                        "schema": {
                            "$ref": "#/definitions/DeleteAllRowsResponse"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful delete-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/commands/delete/{id}": {
            "delete": {
                "summary": "Delete row with given id in table Commands",
                "tags": [
                    "DataBase"
                ],
                "operationId": "DeleteRow",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "description": "Row's identifier",
                        "example": "1"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful delete-request",
                        "schema": {
                            "$ref": "#/definitions/DeleteRowResponse"
                        }
                    },
                    "default": {
                        "description": "Unsuccessful delete-request",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Creation": {
            "type": "object",
            "required": [
                "command"
            ],
            "properties": {
                "command": {
                    "type": "string",
                    "example": "echo hello world"
                }
            }
        },
        "CreationResponse": {
            "type": "object",
            "required": [
                "id",
                "result"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "result": {
                    "type": "string",
                    "example": "hello world\n"
                }
            }
        },
        "RunResponse": {
            "type": "object",
            "required": [
                "result"
            ],
            "properties": {
                "result": {
                    "type": "string",
                    "example": "hello world\n"
                }
            }
        },
        "DeleteAllRowsResponse": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "all rows have been successfully deleted"
                }
            }
        },
        "DeleteRowResponse": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "row 1 has been successfully deleted"
                }
            }
        },
        "Command": {
            "type": "object",
            "required": [
                "id",
                "command",
                "result"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "command": {
                    "type": "string",
                    "example": "echo hello world"
                },
                "result": {
                    "type": "string",
                    "example": "hello world\n"
                }
            }
        },
        "Commands": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/Command"
            }
        },
        "Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "text": {
                    "type": "string",
                    "example": "some internal error was occurred"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BashExecAPI",
	Description:      "This web application provides API for executing bash-scripts.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
