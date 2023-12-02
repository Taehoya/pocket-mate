// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/countries": {
            "get": {
                "description": "get a list of all available countires",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "country"
                ],
                "summary": "Get all country",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CountryResponseDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/healthcheck": {
            "get": {
                "description": "check server status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "check server status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/transactions": {
            "post": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "register transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "register transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "register transaction",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/transactions/options": {
            "get": {
                "description": "get transaction option",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "get transaction option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.TransactionOption"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/transactions/{id}": {
            "put": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "update transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "update transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update transaction",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "delete transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "delete transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/trips": {
            "get": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "get trip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trip"
                ],
                "summary": "get trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TripStatusResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "register trip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trip"
                ],
                "summary": "register trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "register trip",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TripRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/trips/options": {
            "get": {
                "description": "get trip option",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trip"
                ],
                "summary": "get trip option",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.TripNoteOptions"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/trips/{id}": {
            "put": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "update trip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trip"
                ],
                "summary": "update trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update trip",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TripRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearer": []
                    }
                ],
                "description": "delete trip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trip"
                ],
                "summary": "delete trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "User registration",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/v1/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TokenDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BaseResponseDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "dto.CountryResponseDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "CA"
                },
                "currency": {
                    "type": "string",
                    "example": "$"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Canada"
                }
            }
        },
        "dto.ErrorResponseDTO": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "dto.TokenDTO": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "abc.abc.abc"
                },
                "token_type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "dto.TransactionOption": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name_en": {
                    "type": "string",
                    "example": "Food"
                },
                "name_ko": {
                    "type": "string",
                    "example": "음식"
                }
            }
        },
        "dto.TransactionRequestDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 2000.12
                },
                "categoryId": {
                    "type": "integer",
                    "example": 1
                },
                "description": {
                    "type": "string",
                    "example": "sample-description"
                },
                "name": {
                    "type": "string",
                    "example": "sample-name"
                },
                "transactionDateTime": {
                    "type": "string",
                    "example": "2023-11-25T15:04:05Z"
                },
                "tripId": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.TripNoteOptions": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "default": 0,
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "default": "spiralbound",
                    "example": "spiralbound"
                }
            }
        },
        "dto.TripNoteProperty": {
            "type": "object",
            "properties": {
                "boundColor": {
                    "type": "string",
                    "example": "#111111"
                },
                "boundId": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/entities.Bound"
                        }
                    ],
                    "example": 1
                },
                "noteColor": {
                    "type": "string",
                    "example": "#000000"
                }
            }
        },
        "dto.TripRequestDTO": {
            "type": "object",
            "required": [
                "budget",
                "countryId",
                "description",
                "endDateTime",
                "name",
                "noteProperty",
                "startDateTime"
            ],
            "properties": {
                "budget": {
                    "type": "number",
                    "example": 2000.12
                },
                "countryId": {
                    "type": "integer",
                    "example": 1
                },
                "description": {
                    "type": "string",
                    "example": "sample-description"
                },
                "endDateTime": {
                    "type": "string",
                    "example": "2024-01-02T15:04:05Z"
                },
                "name": {
                    "type": "string",
                    "example": "sample-name"
                },
                "noteProperty": {
                    "$ref": "#/definitions/dto.TripNoteProperty"
                },
                "startDateTime": {
                    "type": "string",
                    "example": "2024-01-02T15:04:05Z"
                }
            }
        },
        "dto.TripResponseDTO": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "number",
                    "example": 2000.12
                },
                "countryId": {
                    "type": "integer",
                    "example": 1
                },
                "description": {
                    "type": "string",
                    "example": "sample-description"
                },
                "endDateTime": {
                    "type": "string",
                    "example": "2024-01-02T15:04:05Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "sample-name"
                },
                "noteProperty": {
                    "$ref": "#/definitions/dto.TripNoteProperty"
                },
                "startDateTime": {
                    "type": "string",
                    "example": "2024-01-02T15:04:05Z"
                }
            }
        },
        "dto.TripStatusResponseDTO": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TripResponseDTO"
                    }
                },
                "future": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TripResponseDTO"
                    }
                },
                "past": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TripResponseDTO"
                    }
                }
            }
        },
        "dto.UserRequestDTO": {
            "type": "object",
            "required": [
                "Email",
                "Password"
            ],
            "properties": {
                "Email": {
                    "type": "string",
                    "example": "test@email.com"
                },
                "Password": {
                    "type": "string",
                    "example": "test-password"
                }
            }
        },
        "entities.Bound": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "SpiralBound",
                "GlueBound"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Pocket Mate API",
	Description:      "This is a pocket-mate server api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
