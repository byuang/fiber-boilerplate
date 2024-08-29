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
        "/customers": {
            "get": {
                "description": "Get all customers.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Get all customers.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "true/false",
                        "name": "all",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "start_date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end_date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.CustomerResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create customer",
                "parameters": [
                    {
                        "type": "string",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonCreated"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            }
        },
        "/customers/batch": {
            "post": {
                "description": "Create customer batch.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create customer batch",
                "parameters": [
                    {
                        "description": "create customer batch",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCustomerBatchRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonCreated"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete batch customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Delete batch customer",
                "parameters": [
                    {
                        "description": "delete batch customer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteBatchCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            }
        },
        "/customers/export": {
            "get": {
                "description": "Export Excel customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Export Excel customer.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "true",
                        "name": "all",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "start_date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end_date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            }
        },
        "/customers/import": {
            "post": {
                "description": "Import Excel customer.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Import Excel customer.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Import Excel customer",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            }
        },
        "/customers/{customerId}": {
            "get": {
                "description": "get customer by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "get customer by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer_id",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.CustomerResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            },
            "patch": {
                "description": "update customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "update customer",
                "parameters": [
                    {
                        "description": "update customer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCustomerRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "customer_id",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JsonSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonBadRequest"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.JsonInternalServerError"
                        }
                    }
                }
            }
        },
        "/vehicles": {
            "get": {
                "description": "Get All vehicles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle"
                ],
                "summary": "Get All vehicles.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "is_active",
                        "name": "is_active",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.VehicleResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateCustomerBatchRequest": {
            "type": "object",
            "required": [
                "customers"
            ],
            "properties": {
                "customers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateCustomerRequest"
                    }
                }
            }
        },
        "dto.CreateCustomerRequest": {
            "type": "object",
            "required": [
                "address",
                "email",
                "phone",
                "username"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CustomerResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.DeleteBatchCustomerRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "dto.JsonBadRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "email": "email is required",
                        "username": "username is required"
                    }
                },
                "status": {
                    "type": "string",
                    "example": "BAD REQUEST"
                },
                "trace_id": {
                    "type": "string",
                    "example": "dedc5250-5c20-48c9-9383-fac3ccff2679"
                }
            }
        },
        "dto.JsonCreated": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 201
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Created"
                },
                "status": {
                    "type": "string",
                    "example": "CREATED"
                },
                "trace_id": {
                    "type": "string",
                    "example": "dedc5250-5c20-48c9-9383-fac3ccff2679"
                }
            }
        },
        "dto.JsonInternalServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "errors": {
                    "type": "string",
                    "example": "error database or third party"
                },
                "status": {
                    "type": "string",
                    "example": "INTERNAL SERVER ERROR"
                },
                "trace_id": {
                    "type": "string",
                    "example": "dedc5250-5c20-48c9-9383-fac3ccff2679"
                }
            }
        },
        "dto.JsonNotFound": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "errors": {
                    "type": "string",
                    "example": "record not found"
                },
                "status": {
                    "type": "string",
                    "example": "NOT FOUND"
                },
                "trace_id": {
                    "type": "string",
                    "example": "dedc5250-5c20-48c9-9383-fac3ccff2679"
                }
            }
        },
        "dto.JsonSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "trace_id": {
                    "type": "string",
                    "example": "dedc5250-5c20-48c9-9383-fac3ccff2679"
                }
            }
        },
        "dto.Meta": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_data": {
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "dto.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/dto.Meta"
                },
                "status": {
                    "type": "string"
                },
                "trace_id": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateCustomerRequest": {
            "type": "object",
            "required": [
                "address",
                "email",
                "id",
                "phone",
                "username"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.VehicleResponse": {
            "type": "object",
            "properties": {
                "driver_id": {
                    "type": "integer"
                },
                "driver_name": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "helper_id": {
                    "type": "integer"
                },
                "helper_name": {
                    "type": "string"
                },
                "length": {
                    "type": "number"
                },
                "vehicle_desc": {
                    "type": "string"
                },
                "vehicle_id": {
                    "type": "integer"
                },
                "vehicle_no": {
                    "type": "string"
                },
                "vehicle_type_name": {
                    "type": "string"
                },
                "volume": {
                    "type": "number"
                },
                "width": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
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
	Title:            "Boilerplate API",
	Description:      "Boilerplate API in Go using Fiber framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
