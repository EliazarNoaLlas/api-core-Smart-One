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
        "/api/v1/core/economic_activities/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get economic activities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EconomicActivities"
                ],
                "summary": "get economic activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the cuui id",
                        "name": "cuui_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "the description of the economic activities",
                        "name": "description",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.economicActivitiesResult"
                        }
                    },
                    "500": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.SmartError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.EconomicActivity": {
            "type": "object",
            "required": [
                "created_at",
                "cuui_id",
                "id",
                "status"
            ],
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-12-04 16:01:51"
                },
                "cuui_id": {
                    "type": "string",
                    "example": "0111"
                },
                "description": {
                    "type": "string",
                    "example": "CULTIVO DE ARROZ"
                },
                "id": {
                    "type": "string",
                    "example": "70402269-92be-11ee-a040-0242ac11000e"
                },
                "status": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "domain.PaginationResults": {
            "type": "object",
            "required": [
                "current_page",
                "last_page",
                "size_page",
                "total"
            ],
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "from": {
                    "type": "integer"
                },
                "last_page": {
                    "type": "integer"
                },
                "size_page": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "errorDomain.LayerErr": {
            "type": "string",
            "enum": [
                "domain",
                "infrastructure",
                "interface",
                "use_case"
            ],
            "x-enum-varnames": [
                "Domain",
                "Infra",
                "Interface",
                "UseCase"
            ]
        },
        "errorDomain.LevelErr": {
            "type": "string",
            "enum": [
                "info",
                "warning",
                "error",
                "fatal"
            ],
            "x-enum-varnames": [
                "LevelInfo",
                "LevelWarning",
                "LevelError",
                "LevelFatal"
            ]
        },
        "errorDomain.SmartError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "error": {},
                "function": {
                    "type": "string"
                },
                "httpStatus": {
                    "type": "integer"
                },
                "layer": {
                    "$ref": "#/definitions/errorDomain.LayerErr"
                },
                "level": {
                    "$ref": "#/definitions/errorDomain.LevelErr"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "raw": {
                    "type": "string"
                }
            }
        },
        "rest.economicActivitiesResult": {
            "type": "object",
            "required": [
                "data",
                "pagination",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.EconomicActivity"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/domain.PaginationResults"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
