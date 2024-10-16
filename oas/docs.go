// Package oas Code generated by swaggo/swag. DO NOT EDIT
package oas

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/booking": {
            "post": {
                "description": "Books a fly given a user and the trip information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UseCases"
                ],
                "summary": "Books a fly",
                "parameters": [
                    {
                        "description": "The user fly trip reservation",
                        "name": "trip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/useCases.BookingFlyDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/trips": {
            "post": {
                "description": "creates an entity TripEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "create a trip",
                "parameters": [
                    {
                        "description": "The user fly trip reservations",
                        "name": "trip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/trip.PartialDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/trips/{id}": {
            "get": {
                "description": "get an entity TripEntity by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "get a trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "trip Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "updates an entity TripEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "update a trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "trip Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user fly trip reservations",
                        "name": "trip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/trip.UpdatableDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove an entity TripEntity by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "remove a trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "trip Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "creates an entity UserEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "description": "The user system",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PartialDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get an entity UserEntity by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.FullDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "updates an entity UserEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user system",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdatableDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove an entity UserEntity by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "remove a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}/trips": {
            "get": {
                "description": "get list of entity TripEntity given a userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "get a trip list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "the pagination page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "the amount of items per page",
                        "name": "items",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/trip.FullDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "creates an entity TripEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "create a trip",
                "parameters": [
                    {
                        "description": "The user fly trip reservations",
                        "name": "trip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/trip.PartialDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove all entity TripEntity by userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "remove all trip given a userId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}/trips/{tripId}": {
            "put": {
                "description": "updates an entity TripEntity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "update a trip",
                "parameters": [
                    {
                        "description": "The user fly trip reservations",
                        "name": "trip",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/trip.UpdatableDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/trip.FullDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove an entity TripEntity by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trips"
                ],
                "summary": "remove a trip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "trip Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "trip.FullDTO": {
            "type": "object",
            "properties": {
                "arrival": {
                    "type": "string"
                },
                "created": {
                    "type": "integer"
                },
                "departure": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "miles": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "updated": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "trip.PartialDTO": {
            "type": "object",
            "properties": {
                "arrival": {
                    "type": "string"
                },
                "departure": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "miles": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "trip.UpdatableDTO": {
            "type": "object",
            "properties": {
                "arrival": {
                    "type": "string"
                },
                "departure": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "miles": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "useCases.BookingFlyDTO": {
            "type": "object",
            "properties": {
                "arrival": {
                    "type": "string"
                },
                "departure": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "miles": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "user.FullDTO": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "miles": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "plan": {
                    "type": "string"
                },
                "trips": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/trip.FullDTO"
                    }
                },
                "updated": {
                    "type": "integer"
                }
            }
        },
        "user.PartialDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "miles": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "plan": {
                    "type": "string"
                }
            }
        },
        "user.UpdatableDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "miles": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "plan": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Booking Fly Example API",
	Description:      "This is a sample server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
