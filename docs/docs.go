// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ClearCloud Team",
            "url": "https://github.com/ChappIO/clearcloud/issues",
            "email": "thomas.biesaart@protonmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/libraries": {
            "get": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Search the collection of libraries",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "The page number to fetch",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 20,
                        "description": "The maximum number of elements to return",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.LibraryPage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Add a new library",
                "parameters": [
                    {
                        "description": "The new library",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateLibraryDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Library"
                        }
                    }
                }
            }
        },
        "/api/libraries/{libraryId}": {
            "get": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Get library details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The library id",
                        "name": "libraryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.LibraryDetails"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "description": "This does not delete the files in the library from the disk",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Delete a library",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The library id",
                        "name": "libraryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Update an existing library",
                "parameters": [
                    {
                        "description": "The changes",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateLibraryDTO"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "The library id",
                        "name": "libraryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Library"
                        }
                    }
                }
            }
        },
        "/api/libraries/{libraryId}/shares": {
            "post": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Share a library with a user",
                "parameters": [
                    {
                        "description": "The rights to grant to a specific user",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ShareLibraryDTO"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "The library id",
                        "name": "libraryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/api/libraries/{libraryId}/shares/{userId}": {
            "delete": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Unshare a library",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The library id",
                        "name": "libraryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The user id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get the currently authenticated user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Update the currently authenticated user",
                "parameters": [
                    {
                        "description": "The changes",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Search the collection of users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "The page number to fetch",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 20,
                        "description": "The maximum number of elements to return",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.UserPage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Add a new user",
                "parameters": [
                    {
                        "description": "The new user",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "409": {
                        "description": "This username is already taken",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/api/users/{userId}": {
            "get": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get user details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "OAuth2": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "description": "The changes",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateUserDTO"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "The user id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ApiError": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "details": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "controller.CreateLibraryDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "generic",
                        "books",
                        "movies",
                        "shows",
                        "music"
                    ]
                }
            }
        },
        "controller.CreateUserDTO": {
            "type": "object",
            "required": [
                "firstName",
                "lastName",
                "password",
                "username"
            ],
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.LibraryDetails": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rootFolder": {
                    "type": "string"
                },
                "sharedWith": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CanAccessLibrary"
                    }
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "generic",
                        "books",
                        "movies",
                        "shows",
                        "music"
                    ]
                }
            }
        },
        "controller.LibraryPage": {
            "type": "object",
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.LibrarySummary"
                    }
                },
                "totalElements": {
                    "type": "integer"
                }
            }
        },
        "controller.LibrarySummary": {
            "type": "object",
            "properties": {
                "canWrite": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "generic",
                        "books",
                        "movies",
                        "shows",
                        "music"
                    ]
                }
            }
        },
        "controller.ShareLibraryDTO": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "canWrite": {
                    "type": "boolean"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "controller.UpdateLibraryDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.UserPage": {
            "type": "object",
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.UserSummary"
                    }
                },
                "totalElements": {
                    "type": "integer"
                }
            }
        },
        "controller.UserSummary": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.CanAccessLibrary": {
            "type": "object",
            "properties": {
                "canWrite": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        },
        "model.Library": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rootFolder": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "generic",
                        "books",
                        "movies",
                        "shows",
                        "music"
                    ]
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/oauth2/token"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "development",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "ClearCloud API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
