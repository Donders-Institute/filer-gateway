// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "filer gateway APIs",
    "title": "filer-gateway",
    "version": "0.1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/ping": {
      "get": {
        "security": [
          {
            "oauth2": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "endpoint for API server health check.",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "string",
              "enum": [
                "pong"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/projects": {
      "post": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "provision filer resource for a new project.",
        "parameters": [
          {
            "description": "data for project provisioning",
            "name": "projectProvisionData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyProjectProvision"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/projects/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "get filer resource for an existing project.",
        "parameters": [
          {
            "type": "string",
            "description": "project identifier",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyProjectResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "project not found",
            "schema": {
              "type": "string",
              "enum": [
                "project not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "update filer resource for an existing project.",
        "parameters": [
          {
            "type": "string",
            "description": "project identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "data for project update",
            "name": "projectUpdateData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyProjectResource"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "project not found",
            "schema": {
              "type": "string",
              "enum": [
                "project not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/tasks/{type}/{id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "query background task status",
        "parameters": [
          {
            "type": "string",
            "description": "task identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "task type",
            "name": "type",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "provision filer resource for a new user.",
        "parameters": [
          {
            "description": "data for user provisioning",
            "name": "userProvisionData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyUserProvision"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "get filer resource for an existing user.",
        "parameters": [
          {
            "type": "string",
            "description": "user identifier",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyUserResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "user not found",
            "schema": {
              "type": "string",
              "enum": [
                "user not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "update filer resource for an existing user.",
        "parameters": [
          {
            "type": "string",
            "description": "user identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "data for user update",
            "name": "userUpdateData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyUserResource"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "user not found",
            "schema": {
              "type": "string",
              "enum": [
                "user not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "member": {
      "description": "JSON object for a project member.",
      "required": [
        "userID",
        "role"
      ],
      "properties": {
        "role": {
          "description": "role of the member. Use the value \"none\" in request data to remove user from the project membership.",
          "type": "string",
          "enum": [
            "manager",
            "contributor",
            "viewer",
            "traverse",
            "none"
          ]
        },
        "userID": {
          "description": "userid of the member.",
          "type": "string"
        }
      }
    },
    "memberOf": {
      "description": "a list providing an overview of user's member roles in all active projects.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/projectRole"
      }
    },
    "members": {
      "description": "a list of project members.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/member"
      }
    },
    "principle": {
      "type": "string"
    },
    "projectID": {
      "description": "project identifier.",
      "type": "string"
    },
    "projectRole": {
      "description": "JSON object for a member role of a project.",
      "required": [
        "projectID",
        "role"
      ],
      "properties": {
        "projectID": {
          "description": "project identifier",
          "type": "string"
        },
        "role": {
          "description": "role of the member.",
          "type": "string",
          "enum": [
            "manager",
            "contributor",
            "viewer",
            "traverse"
          ]
        }
      }
    },
    "requestBodyProjectProvision": {
      "required": [
        "projectID",
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "requestBodyProjectResource": {
      "description": "JSON object describing resource to be set to the project.",
      "required": [
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "requestBodyUserProvision": {
      "required": [
        "userID",
        "storage"
      ],
      "properties": {
        "storage": {
          "$ref": "#/definitions/storageRequest"
        },
        "userID": {
          "$ref": "#/definitions/userID"
        }
      }
    },
    "requestBodyUserResource": {
      "description": "JSON object describing resource to be set to the user.",
      "required": [
        "storage"
      ],
      "properties": {
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "responseBody400": {
      "description": "JSON object containing error message concerning bad client request.",
      "properties": {
        "errorMessage": {
          "description": "error message specifying the bad request.",
          "type": "string"
        }
      }
    },
    "responseBody500": {
      "description": "JSON object containing server side error.",
      "properties": {
        "errorMessage": {
          "description": "server-side error message.",
          "type": "string"
        },
        "exitCode": {
          "description": "server-side exit code.",
          "type": "integer"
        }
      }
    },
    "responseBodyProjectMembers": {
      "description": "JSON object containing current list of members implemented on the filer.",
      "required": [
        "projectID",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        }
      }
    },
    "responseBodyProjectResource": {
      "description": "JSON object containing project resources.",
      "required": [
        "projectID",
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        }
      }
    },
    "responseBodyProjectStorage": {
      "description": "JSON object containing storage resource information of a project.",
      "required": [
        "projectID",
        "storage"
      ],
      "properties": {
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        }
      }
    },
    "responseBodyTaskResource": {
      "description": "JSON object containing scheduled task resource.",
      "required": [
        "taskID",
        "taskStatus"
      ],
      "properties": {
        "taskID": {
          "$ref": "#/definitions/taskID"
        },
        "taskStatus": {
          "$ref": "#/definitions/taskStatus"
        }
      }
    },
    "responseBodyUserResource": {
      "description": "JSON object containing user resources.",
      "required": [
        "userID",
        "memberOf",
        "storage"
      ],
      "properties": {
        "memberOf": {
          "$ref": "#/definitions/memberOf"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        },
        "userID": {
          "$ref": "#/definitions/userID"
        }
      }
    },
    "storageRequest": {
      "description": "JSON object for storage resource data.",
      "required": [
        "system",
        "quotaGb"
      ],
      "properties": {
        "quotaGb": {
          "description": "storage quota in GiB.",
          "type": "integer"
        },
        "system": {
          "description": "the targeting filer system on which the storage resource is allocated. Use the value \"none\" to skip the storage resource setting.",
          "type": "string",
          "enum": [
            "netapp",
            "freenas",
            "cephfs",
            "none"
          ]
        }
      }
    },
    "storageResponse": {
      "description": "JSON object for storage resource data.",
      "required": [
        "system",
        "quotaGb",
        "usageMb"
      ],
      "properties": {
        "quotaGb": {
          "description": "assigned storage quota in GiB.",
          "type": "integer"
        },
        "system": {
          "description": "the targeting filer on which the storage resource is allocated.",
          "type": "string",
          "enum": [
            "netapp",
            "freenas",
            "cephfs"
          ]
        },
        "usageMb": {
          "description": "used storage quota in MiB.",
          "type": "integer"
        }
      }
    },
    "taskID": {
      "description": "identifier for scheduled background tasks.",
      "type": "string"
    },
    "taskStatus": {
      "description": "status of the background task.",
      "required": [
        "status",
        "error",
        "result"
      ],
      "properties": {
        "error": {
          "description": "task error message from the last execution.",
          "type": "string"
        },
        "result": {
          "description": "task result from the last execution.",
          "type": "string"
        },
        "status": {
          "description": "task status from the last execution.",
          "type": "string",
          "enum": [
            "waiting",
            "processing",
            "failed",
            "succeeded",
            "canceled"
          ]
        }
      }
    },
    "userID": {
      "description": "user identifier.",
      "type": "string"
    }
  },
  "securityDefinitions": {
    "apiKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    },
    "basicAuth": {
      "type": "basic"
    },
    "oauth2": {
      "type": "oauth2",
      "flow": "application",
      "tokenUrl": "https://auth-dev.dccn.nl/connect/token",
      "scopes": {
        "project-database-core-api": "general access scope for filer-gateway APIs"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "filer gateway APIs",
    "title": "filer-gateway",
    "version": "0.1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/ping": {
      "get": {
        "security": [
          {
            "oauth2": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "endpoint for API server health check.",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "string",
              "enum": [
                "pong"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/projects": {
      "post": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "provision filer resource for a new project.",
        "parameters": [
          {
            "description": "data for project provisioning",
            "name": "projectProvisionData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyProjectProvision"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/projects/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "get filer resource for an existing project.",
        "parameters": [
          {
            "type": "string",
            "description": "project identifier",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyProjectResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "project not found",
            "schema": {
              "type": "string",
              "enum": [
                "project not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "update filer resource for an existing project.",
        "parameters": [
          {
            "type": "string",
            "description": "project identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "data for project update",
            "name": "projectUpdateData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyProjectResource"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "project not found",
            "schema": {
              "type": "string",
              "enum": [
                "project not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/tasks/{type}/{id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "query background task status",
        "parameters": [
          {
            "type": "string",
            "description": "task identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "task type",
            "name": "type",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "provision filer resource for a new user.",
        "parameters": [
          {
            "description": "data for user provisioning",
            "name": "userProvisionData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyUserProvision"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "get filer resource for an existing user.",
        "parameters": [
          {
            "type": "string",
            "description": "user identifier",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyUserResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "user not found",
            "schema": {
              "type": "string",
              "enum": [
                "user not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "apiKeyHeader": [],
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "update filer resource for an existing user.",
        "parameters": [
          {
            "type": "string",
            "description": "user identifier",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "data for user update",
            "name": "userUpdateData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestBodyUserResource"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/responseBodyTaskResource"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/responseBody400"
            }
          },
          "404": {
            "description": "user not found",
            "schema": {
              "type": "string",
              "enum": [
                "user not found"
              ]
            }
          },
          "500": {
            "description": "failure",
            "schema": {
              "$ref": "#/definitions/responseBody500"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "member": {
      "description": "JSON object for a project member.",
      "required": [
        "userID",
        "role"
      ],
      "properties": {
        "role": {
          "description": "role of the member. Use the value \"none\" in request data to remove user from the project membership.",
          "type": "string",
          "enum": [
            "manager",
            "contributor",
            "viewer",
            "traverse",
            "none"
          ]
        },
        "userID": {
          "description": "userid of the member.",
          "type": "string"
        }
      }
    },
    "memberOf": {
      "description": "a list providing an overview of user's member roles in all active projects.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/projectRole"
      }
    },
    "members": {
      "description": "a list of project members.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/member"
      }
    },
    "principle": {
      "type": "string"
    },
    "projectID": {
      "description": "project identifier.",
      "type": "string"
    },
    "projectRole": {
      "description": "JSON object for a member role of a project.",
      "required": [
        "projectID",
        "role"
      ],
      "properties": {
        "projectID": {
          "description": "project identifier",
          "type": "string"
        },
        "role": {
          "description": "role of the member.",
          "type": "string",
          "enum": [
            "manager",
            "contributor",
            "viewer",
            "traverse"
          ]
        }
      }
    },
    "requestBodyProjectProvision": {
      "required": [
        "projectID",
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "requestBodyProjectResource": {
      "description": "JSON object describing resource to be set to the project.",
      "required": [
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "requestBodyUserProvision": {
      "required": [
        "userID",
        "storage"
      ],
      "properties": {
        "storage": {
          "$ref": "#/definitions/storageRequest"
        },
        "userID": {
          "$ref": "#/definitions/userID"
        }
      }
    },
    "requestBodyUserResource": {
      "description": "JSON object describing resource to be set to the user.",
      "required": [
        "storage"
      ],
      "properties": {
        "storage": {
          "$ref": "#/definitions/storageRequest"
        }
      }
    },
    "responseBody400": {
      "description": "JSON object containing error message concerning bad client request.",
      "properties": {
        "errorMessage": {
          "description": "error message specifying the bad request.",
          "type": "string"
        }
      }
    },
    "responseBody500": {
      "description": "JSON object containing server side error.",
      "properties": {
        "errorMessage": {
          "description": "server-side error message.",
          "type": "string"
        },
        "exitCode": {
          "description": "server-side exit code.",
          "type": "integer"
        }
      }
    },
    "responseBodyProjectMembers": {
      "description": "JSON object containing current list of members implemented on the filer.",
      "required": [
        "projectID",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        }
      }
    },
    "responseBodyProjectResource": {
      "description": "JSON object containing project resources.",
      "required": [
        "projectID",
        "storage",
        "members"
      ],
      "properties": {
        "members": {
          "$ref": "#/definitions/members"
        },
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        }
      }
    },
    "responseBodyProjectStorage": {
      "description": "JSON object containing storage resource information of a project.",
      "required": [
        "projectID",
        "storage"
      ],
      "properties": {
        "projectID": {
          "$ref": "#/definitions/projectID"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        }
      }
    },
    "responseBodyTaskResource": {
      "description": "JSON object containing scheduled task resource.",
      "required": [
        "taskID",
        "taskStatus"
      ],
      "properties": {
        "taskID": {
          "$ref": "#/definitions/taskID"
        },
        "taskStatus": {
          "$ref": "#/definitions/taskStatus"
        }
      }
    },
    "responseBodyUserResource": {
      "description": "JSON object containing user resources.",
      "required": [
        "userID",
        "memberOf",
        "storage"
      ],
      "properties": {
        "memberOf": {
          "$ref": "#/definitions/memberOf"
        },
        "storage": {
          "$ref": "#/definitions/storageResponse"
        },
        "userID": {
          "$ref": "#/definitions/userID"
        }
      }
    },
    "storageRequest": {
      "description": "JSON object for storage resource data.",
      "required": [
        "system",
        "quotaGb"
      ],
      "properties": {
        "quotaGb": {
          "description": "storage quota in GiB.",
          "type": "integer"
        },
        "system": {
          "description": "the targeting filer system on which the storage resource is allocated. Use the value \"none\" to skip the storage resource setting.",
          "type": "string",
          "enum": [
            "netapp",
            "freenas",
            "cephfs",
            "none"
          ]
        }
      }
    },
    "storageResponse": {
      "description": "JSON object for storage resource data.",
      "required": [
        "system",
        "quotaGb",
        "usageMb"
      ],
      "properties": {
        "quotaGb": {
          "description": "assigned storage quota in GiB.",
          "type": "integer"
        },
        "system": {
          "description": "the targeting filer on which the storage resource is allocated.",
          "type": "string",
          "enum": [
            "netapp",
            "freenas",
            "cephfs"
          ]
        },
        "usageMb": {
          "description": "used storage quota in MiB.",
          "type": "integer"
        }
      }
    },
    "taskID": {
      "description": "identifier for scheduled background tasks.",
      "type": "string"
    },
    "taskStatus": {
      "description": "status of the background task.",
      "required": [
        "status",
        "error",
        "result"
      ],
      "properties": {
        "error": {
          "description": "task error message from the last execution.",
          "type": "string"
        },
        "result": {
          "description": "task result from the last execution.",
          "type": "string"
        },
        "status": {
          "description": "task status from the last execution.",
          "type": "string",
          "enum": [
            "waiting",
            "processing",
            "failed",
            "succeeded",
            "canceled"
          ]
        }
      }
    },
    "userID": {
      "description": "user identifier.",
      "type": "string"
    }
  },
  "securityDefinitions": {
    "apiKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    },
    "basicAuth": {
      "type": "basic"
    },
    "oauth2": {
      "type": "oauth2",
      "flow": "application",
      "tokenUrl": "https://auth-dev.dccn.nl/connect/token",
      "scopes": {
        "project-database-core-api": "general access scope for filer-gateway APIs"
      }
    }
  }
}`))
}
