// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

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
        "/packages": {
            "get": {
                "description": "Returns user` + "`" + `s packages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Get packages",
                "operationId": "getPackages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/v1.PackageResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Crates a package",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Create",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PackageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.PackageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/packages/{packageID}": {
            "get": {
                "description": "Returns package by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Get package",
                "operationId": "getPackage",
                "parameters": [
                    {
                        "type": "string",
                        "example": "6155c774-d1e2-4816-b7f4-52ebb949f044",
                        "description": "package ID",
                        "name": "packageID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.PackageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates a package",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "packages"
                ],
                "summary": "Update",
                "operationId": "updatePackage",
                "parameters": [
                    {
                        "type": "string",
                        "example": "6155c774-d1e2-4816-b7f4-52ebb949f044",
                        "description": "package ID",
                        "name": "packageID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PackageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.PackageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/users/user/auth": {
            "post": {
                "description": "Authenticates user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Auth",
                "operationId": "auth",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.authUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/users/user/refresh": {
            "post": {
                "description": "Refreshes users JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Refresh",
                "operationId": "refresh",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.refreshAuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/users/user/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "operationId": "register",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.registerUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/users/user/update": {
            "patch": {
                "description": "Updates users data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update",
                "operationId": "update",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.updateUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.PackageRequest": {
            "type": "object",
            "required": [
                "height",
                "name",
                "weight",
                "width"
            ],
            "properties": {
                "height": {
                    "type": "number",
                    "example": 15
                },
                "name": {
                    "type": "string",
                    "example": "Package for Moxem"
                },
                "weight": {
                    "type": "number",
                    "example": 11.3
                },
                "width": {
                    "type": "number",
                    "example": 13.8
                }
            }
        },
        "v1.PackageResponse": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "number",
                    "example": 15
                },
                "name": {
                    "type": "string",
                    "example": "Package for Moxem"
                },
                "ownerID": {
                    "type": "string",
                    "example": "P1873eecd-c2d0-4aa2-a8d4-e0de232c5ac6"
                },
                "packageID": {
                    "type": "string",
                    "example": "6155c774-d1e2-4816-b7f4-52ebb949f044"
                },
                "status": {
                    "type": "string",
                    "example": "new"
                },
                "weight": {
                    "type": "number",
                    "example": 11.3
                },
                "width": {
                    "type": "number",
                    "example": 13.8
                }
            }
        },
        "v1.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "v1.authUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "qwerty123"
                }
            }
        },
        "v1.refreshAuthRequest": {
            "type": "object",
            "required": [
                "refresh_token",
                "user_id"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "v1.registerUserRequest": {
            "type": "object",
            "required": [
                "deliveryAddress",
                "email",
                "lastName",
                "name",
                "password"
            ],
            "properties": {
                "deliveryAddress": {
                    "type": "string",
                    "example": "Pushkina street"
                },
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "qwerty123"
                }
            }
        },
        "v1.registerUserResponse": {
            "type": "object",
            "required": [
                "deliveryAddress",
                "email",
                "id",
                "lastName",
                "name"
            ],
            "properties": {
                "deliveryAddress": {
                    "type": "string",
                    "example": "Pushkina street"
                },
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "id": {
                    "type": "string",
                    "example": "d9e48656-ae36-4fde-af78-5f6250e11ead"
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                }
            }
        },
        "v1.updateUserRequest": {
            "type": "object",
            "required": [
                "deliveryAddress",
                "lastName",
                "name",
                "password"
            ],
            "properties": {
                "deliveryAddress": {
                    "type": "string",
                    "example": "Pushkina street"
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "qwerty123"
                }
            }
        },
        "v1.updateUserResponse": {
            "type": "object",
            "required": [
                "deliveryAddress",
                "email",
                "id",
                "lastName",
                "name"
            ],
            "properties": {
                "deliveryAddress": {
                    "type": "string",
                    "example": "Pushkina street"
                },
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "id": {
                    "type": "string",
                    "example": "d9e48656-ae36-4fde-af78-5f6250e11ead"
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Delivery service API",
	Description:      "Delivery service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
