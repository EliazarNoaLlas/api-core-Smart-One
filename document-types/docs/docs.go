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
        "/api/v1/core/document_types/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get document types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DocumentTypes"
                ],
                "summary": "get document types",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the number of the document type",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "the description of the document type",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "the abbreviated description of the document type",
                        "name": "abbreviated_description",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.documentTypeResult"
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
        "/api/v1/core/document_types/create_document_types/{documentTypeId}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a document type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DocumentTypes"
                ],
                "summary": "Create a document type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "document type id",
                        "name": "documentTypeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create document type body",
                        "name": "createDocumentTypeBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateDocumentTypeBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
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
        "/api/v1/core/document_types/delete_document_types/{documentTypeId}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a document type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DocumentTypes"
                ],
                "summary": "Delete a document type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "document type id",
                        "name": "documentTypeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.deleteDocumentTypesResult"
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
        "/api/v1/core/document_types/update_document_types/{documentTypeId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a document type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DocumentTypes"
                ],
                "summary": "Update a document type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "document type id",
                        "name": "documentTypeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update document types body",
                        "name": "updateDocumentTypeBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateDocumentTypeBody"
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
            }
        }
    },
    "definitions": {
        "domain.CreateDocumentTypeBody": {
            "type": "object",
            "required": [
                "abbreviated_description",
                "description",
                "enable",
                "number"
            ],
            "properties": {
                "abbreviated_description": {
                    "description": "Description: the abbreviation of the type of document",
                    "type": "string",
                    "example": "DNI"
                },
                "description": {
                    "description": "Description: the description of the type of document",
                    "type": "string",
                    "example": "DOCUMENTO NACIONAL DE IDENTIDAD"
                },
                "enable": {
                    "description": "Description: the status of the type of document",
                    "type": "integer",
                    "example": 1
                },
                "number": {
                    "description": "Description: the number of the type of document",
                    "type": "string",
                    "example": "01"
                }
            }
        },
        "domain.DocumentType": {
            "type": "object",
            "required": [
                "abbreviated_description",
                "description",
                "enable",
                "id",
                "number"
            ],
            "properties": {
                "abbreviated_description": {
                    "description": "Description: the abbreviation of the type of document",
                    "type": "string",
                    "example": "DNI"
                },
                "created_at": {
                    "description": "Description: the date of the type of document",
                    "type": "string",
                    "example": "2023-12-05 15:49:56"
                },
                "description": {
                    "description": "Description: the description of the type of document",
                    "type": "string",
                    "example": "DOCUMENTO NACIONAL DE IDENTIDAD"
                },
                "enable": {
                    "description": "Description: the status of the type of document",
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "description": "Description: the id of the type of document",
                    "type": "string",
                    "example": "00a58296-93b4-11ee-a040-0242ac11000e"
                },
                "number": {
                    "description": "Description: the number of the type of document",
                    "type": "string",
                    "example": "01"
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
        "domain.UpdateDocumentTypeBody": {
            "type": "object",
            "required": [
                "abbreviated_description",
                "description",
                "enable",
                "number"
            ],
            "properties": {
                "abbreviated_description": {
                    "description": "Description: the abbreviation of the type of document",
                    "type": "string",
                    "example": "DNI"
                },
                "description": {
                    "description": "Description: the description of the type of document",
                    "type": "string",
                    "example": "DOCUMENTO NACIONAL DE IDENTIDAD"
                },
                "enable": {
                    "description": "Description: the status of the type of document",
                    "type": "integer",
                    "example": 1
                },
                "number": {
                    "description": "Description: the number of the type of document",
                    "type": "string",
                    "example": "01"
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
        "rest.deleteDocumentTypesResult": {
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
        "rest.documentTypeResult": {
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
                        "$ref": "#/definitions/domain.DocumentType"
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
