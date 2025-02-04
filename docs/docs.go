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
        "/urls": {
            "post": {
                "description": "Add a new URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new URL",
                "operationId": "add-url",
                "parameters": [
                    {
                        "description": "URL to be added",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewURL"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.URL"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/urlErrors.GenericError"
                        }
                    }
                }
            }
        },
        "/urls/": {
            "get": {
                "description": "get all available short urls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "get all available short urls",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.URL"
                            }
                        }
                    }
                }
            }
        },
        "/urls/{externalId}/visits": {
            "get": {
                "description": "Count the number of visits for a URL based on its external ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Count URL visits",
                "operationId": "count-url-visits",
                "parameters": [
                    {
                        "type": "string",
                        "description": "External ID of the URL",
                        "name": "externalId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "visits",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/urlErrors.GenericError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/urlErrors.GenericError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.NewURL": {
            "type": "object",
            "required": [
                "longUrl"
            ],
            "properties": {
                "longUrl": {
                    "type": "string"
                }
            }
        },
        "model.URL": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "TODO still as sting on swagger",
                    "type": "string"
                },
                "externalId": {
                    "type": "string"
                },
                "longUrl": {
                    "type": "string"
                }
            }
        },
        "urlErrors.GenericError": {
            "type": "object",
            "properties": {
                "statusCode": {
                    "type": "integer"
                }
            }
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
