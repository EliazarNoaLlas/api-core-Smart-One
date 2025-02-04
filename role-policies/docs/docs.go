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
        "/api/v1/core/roles/{roleId}/policies": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get policies by role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "get policies by role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size page",
                        "name": "size_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.rolePoliciesResult"
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
                "description": "Create role policy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "Create role policy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create role policy body",
                        "name": "createRolePolicyBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateRolePolicyBody"
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
        "/api/v1/core/roles/{roleId}/policies/batch": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create multiple role policies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "Create multiple role policies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create multiple role policy body",
                        "name": "createRolePolicyBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.createMultipleRolePoliciesValidate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/httpResponse.IdsResult"
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
                "description": "Delete multiple role policies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "Delete multiple role policies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.deleteMultipleRolePoliciesValidate"
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
        "/api/v1/core/roles/{roleId}/policies/{rolePolicyId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update role policy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "Update role policy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role policy id",
                        "name": "rolePolicyId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update role policy body",
                        "name": "updateRolePolicyBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateRolePolicyBody"
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
                "description": "Delete role policy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role Policy"
                ],
                "summary": "Delete role policy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role policy id",
                        "name": "rolePolicyId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Request",
                        "schema": {
                            "$ref": "#/definitions/rest.deleteRolePoliciesResult"
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
        "domain.CreateRolePolicyBody": {
            "type": "object",
            "required": [
                "enable",
                "policy_id"
            ],
            "properties": {
                "enable": {
                    "description": "Description: enable of the created role policies",
                    "type": "boolean",
                    "example": true
                },
                "policy_id": {
                    "description": "Description: the policy_id of the created role policies",
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-042hs5278420"
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
        "domain.PolicyByRolePolicy": {
            "type": "object",
            "required": [
                "description",
                "enable",
                "id",
                "level",
                "name"
            ],
            "properties": {
                "created_at": {
                    "description": "Description: the created_at of the role policies",
                    "type": "string",
                    "example": "2023-11-10 08:10:00"
                },
                "description": {
                    "description": "Description: the description of the role policies",
                    "type": "string",
                    "example": "Politica para accesos a logistica requerimientos en todo el conglomerado"
                },
                "enable": {
                    "description": "Description: enable of the role policies",
                    "type": "boolean",
                    "example": true
                },
                "id": {
                    "description": "Description: the id of the role policies",
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-0242ac110016"
                },
                "level": {
                    "description": "Description: the level of the role policies",
                    "type": "string",
                    "example": "system"
                },
                "name": {
                    "description": "Description: the name of the role policies",
                    "type": "string",
                    "example": "LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"
                }
            }
        },
        "domain.RolePolicy": {
            "type": "object",
            "required": [
                "enable",
                "id",
                "policy"
            ],
            "properties": {
                "created_at": {
                    "description": "Description: the created_at of the role policies",
                    "type": "string",
                    "example": "2023-11-10 08:10:00"
                },
                "enable": {
                    "description": "Description: enable of the role policies",
                    "type": "boolean",
                    "example": true
                },
                "id": {
                    "description": "Description: the id of the role policies",
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-0242ac110017"
                },
                "policy": {
                    "$ref": "#/definitions/domain.PolicyByRolePolicy"
                }
            }
        },
        "domain.UpdateRolePolicyBody": {
            "type": "object",
            "required": [
                "enable",
                "policy_id"
            ],
            "properties": {
                "enable": {
                    "description": "Description: enable of the update role policies",
                    "type": "boolean",
                    "example": true
                },
                "policy_id": {
                    "description": "Description: the policy_id of the update role policies",
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-042hs5278420"
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
        "httpResponse.IdsResult": {
            "type": "object",
            "required": [
                "data",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "201"
                    ]
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
        "rest.createMultipleRolePoliciesValidate": {
            "type": "object",
            "properties": {
                "rolePolicies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/rest.createMultipleRolePolicyValidate"
                    }
                }
            }
        },
        "rest.createMultipleRolePolicyValidate": {
            "type": "object",
            "required": [
                "id",
                "policy_id"
            ],
            "properties": {
                "enable": {
                    "type": "boolean",
                    "example": true
                },
                "id": {
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-0442ac210932"
                },
                "policy_id": {
                    "type": "string",
                    "example": "739bbbc9-7e93-11ee-89fd-0442ac210931"
                }
            }
        },
        "rest.deleteMultipleRolePoliciesValidate": {
            "type": "object",
            "required": [
                "role_policy_ids"
            ],
            "properties": {
                "role_policy_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "739bbbc9-7e93-11ee-89fd-0442ac210931"
                    ]
                }
            }
        },
        "rest.deleteRolePoliciesResult": {
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
        "rest.rolePoliciesResult": {
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
                        "$ref": "#/definitions/domain.RolePolicy"
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
