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
        "/api/v1/core/merchants": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Get merchants",
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.merchantsResult"
                        }
                    },
                    "500": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.SmartError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Create merchant",
                "parameters": [
                    {
                        "description": "Create merchant body",
                        "name": "createMerchantBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateMerchantBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/httpResponse.IdResult"
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
        },
        "/api/v1/core/merchants/{merchantId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Update merchant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "merchant id",
                        "name": "merchantId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update merchant body",
                        "name": "updateMerchantBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateMerchantBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/httpResponse.StatusResult"
                        }
                    },
                    "500": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.SmartError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Delete a merchant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "merchant id",
                        "name": "merchantId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.deleteMerchantsResult"
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
        "domain.CreateMerchantBody": {
            "type": "object",
            "required": [
                "address",
                "description",
                "document",
                "image_path",
                "industry",
                "name",
                "phone"
            ],
            "properties": {
                "address": {
                    "description": "Description: the address of the merchant",
                    "type": "string",
                    "example": "123 Main Street"
                },
                "description": {
                    "description": "Description: the description of the merchant",
                    "type": "string",
                    "example": "Proveedor de servicios de mantenimiento"
                },
                "document": {
                    "description": "Description: the document of the merchant",
                    "type": "string",
                    "example": "123456789"
                },
                "image_path": {
                    "description": "Description: the image_path of the merchant",
                    "type": "string",
                    "example": "https://example.com/images/odin_logo.png"
                },
                "industry": {
                    "description": "Description: the industry of the merchant",
                    "type": "string",
                    "example": "Mantenimiento"
                },
                "name": {
                    "description": "Description: the name of the merchant",
                    "type": "string",
                    "example": "Odin Corp"
                },
                "phone": {
                    "description": "Description: the phone of the merchant",
                    "type": "string",
                    "example": "+1234567890"
                }
            }
        },
        "domain.Merchant": {
            "type": "object",
            "required": [
                "address",
                "created_at",
                "description",
                "document",
                "id",
                "image_path",
                "industry",
                "name",
                "phone"
            ],
            "properties": {
                "address": {
                    "description": "Description: the address of the merchant",
                    "type": "string",
                    "example": "123 Main Street"
                },
                "created_at": {
                    "description": "Description: the created_at of the merchant",
                    "type": "string",
                    "example": "2023-11-10 08:10:00"
                },
                "description": {
                    "description": "Description: the description of the merchant",
                    "type": "string",
                    "example": "Proveedor de servicios de mantenimiento"
                },
                "document": {
                    "description": "Description: the document of the merchant",
                    "type": "string",
                    "example": "123456789"
                },
                "id": {
                    "description": "Description: the id of the merchant",
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-0242ac110016"
                },
                "image_path": {
                    "description": "Description: the image_path of the merchant",
                    "type": "string",
                    "example": "https://example.com/images/odin_logo.png"
                },
                "industry": {
                    "description": "Description: the industry of the merchant",
                    "type": "string",
                    "example": "Mantenimiento"
                },
                "name": {
                    "description": "Description: the name of the merchant",
                    "type": "string",
                    "example": "Odin Corp"
                },
                "phone": {
                    "description": "Description: the phone of the merchant",
                    "type": "string",
                    "example": "+1234567890"
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
        "domain.UpdateMerchantBody": {
            "type": "object",
            "required": [
                "address",
                "description",
                "document",
                "image_path",
                "industry",
                "name",
                "phone"
            ],
            "properties": {
                "address": {
                    "description": "Description: the address of the merchant",
                    "type": "string",
                    "example": "123 Main Street"
                },
                "description": {
                    "description": "Description: the description of the merchant",
                    "type": "string",
                    "example": "Proveedor de servicios de mantenimiento"
                },
                "document": {
                    "description": "Description: the document of the merchant",
                    "type": "string",
                    "example": "123456789"
                },
                "image_path": {
                    "description": "Description: the image_path of the merchant",
                    "type": "string",
                    "example": "https://example.com/images/odin_logo.png"
                },
                "industry": {
                    "description": "Description: the industry of the merchant",
                    "type": "string",
                    "example": "Mantenimiento"
                },
                "name": {
                    "description": "Description: the name of the merchant",
                    "type": "string",
                    "example": "Odin Corp"
                },
                "phone": {
                    "description": "Description: the phone of the merchant",
                    "type": "string",
                    "example": "+1234567890"
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
        "httpResponse.IdResult": {
            "type": "object",
            "required": [
                "data",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "string",
                    "example": "201"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "httpResponse.StatusResult": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "rest.deleteMerchantsResult": {
            "type": "object",
            "required": [
                "data",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "boolean"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "rest.merchantsResult": {
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
                        "$ref": "#/definitions/domain.Merchant"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
