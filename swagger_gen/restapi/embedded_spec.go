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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "goliac is a golang REST application skeleton The base path for all the APIs is \"/api/v1\".\n",
    "title": "goliac",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/collaborators": {
      "get": {
        "description": "Get all external collaborators",
        "tags": [
          "app"
        ],
        "operationId": "getCollaborators",
        "responses": {
          "200": {
            "description": "get list of collarators",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/collaborators/{collaboratorID}": {
      "get": {
        "description": "Get collaborator and repos",
        "tags": [
          "app"
        ],
        "operationId": "getCollaborator",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "collaborator name",
            "name": "collaboratorID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get collaborator details especially repositories",
            "schema": {
              "$ref": "#/definitions/collaboratorDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flushcache": {
      "post": {
        "description": "Flush the Github remote cache",
        "tags": [
          "app"
        ],
        "operationId": "postFlushCache",
        "responses": {
          "200": {
            "description": "cache flushed"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/liveness": {
      "get": {
        "description": "Check if Goliac is healthy",
        "tags": [
          "health"
        ],
        "operationId": "getLiveness",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/readiness": {
      "get": {
        "description": "Check if Goliac is ready to serve",
        "tags": [
          "health"
        ],
        "operationId": "getReadiness",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories": {
      "get": {
        "description": "Get all repositories",
        "tags": [
          "app"
        ],
        "operationId": "getRepositories",
        "responses": {
          "200": {
            "description": "get list of repositories",
            "schema": {
              "$ref": "#/definitions/repositories"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{repositoryID}": {
      "get": {
        "description": "Get repository and associated teams",
        "tags": [
          "app"
        ],
        "operationId": "getRepository",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "repository slug name",
            "name": "repositoryID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get repository details especially teams that have access",
            "schema": {
              "$ref": "#/definitions/repositoryDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/resync": {
      "post": {
        "description": "Ask to sync again against Github",
        "tags": [
          "app"
        ],
        "operationId": "postResync",
        "responses": {
          "200": {
            "description": "resync in progress"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Get different statistics on Goliac",
        "tags": [
          "app"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "get Goliac statistics",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/teams": {
      "get": {
        "description": "Get all teams",
        "tags": [
          "app"
        ],
        "operationId": "getTeams",
        "responses": {
          "200": {
            "description": "get list of teams",
            "schema": {
              "$ref": "#/definitions/teams"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/teams/{teamID}": {
      "get": {
        "description": "Get team and associated users and repos",
        "tags": [
          "app"
        ],
        "operationId": "getTeam",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "team name",
            "name": "teamID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get team details especially users and repositories",
            "schema": {
              "$ref": "#/definitions/teamDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "Get all users",
        "tags": [
          "app"
        ],
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "get list of users (organization or external)",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{userID}": {
      "get": {
        "description": "Get user and associated teams and repos",
        "tags": [
          "app"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "user name",
            "name": "userID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get user details especially teams and repositories",
            "schema": {
              "$ref": "#/definitions/userDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "collaboratorDetails": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "repositories": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/repository"
      }
    },
    "repository": {
      "type": "object",
      "properties": {
        "allowUpdateBranch": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "archived": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "autoMergeAllowed": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "deleteBranchOnMerge": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "public": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        }
      }
    },
    "repositoryDetails": {
      "type": "object",
      "properties": {
        "allowUpdateBranch": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "archived": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "autoMergeAllowed": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "collaborators": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "access": {
                "type": "string",
                "minLength": 1
              },
              "name": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        },
        "deleteBranchOnMerge": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "public": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "teams": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "access": {
                "type": "string",
                "minLength": 1
              },
              "name": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        }
      }
    },
    "status": {
      "type": "object",
      "properties": {
        "detailedErrors": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "detailedWarnings": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "lastSyncError": {
          "type": "string"
        },
        "lastSyncTime": {
          "type": "string",
          "minLength": 1
        },
        "nbRepos": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbTeams": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbUsers": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbUsersExternal": {
          "type": "integer",
          "x-omitempty": false
        },
        "version": {
          "type": "string"
        }
      }
    },
    "team": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "owners": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    },
    "teamDetails": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "external": {
                "type": "boolean",
                "x-isnullable": false,
                "x-omitempty": false
              },
              "githubid": {
                "type": "string",
                "x-isnullable": false
              },
              "name": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "owners": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "external": {
                "type": "boolean",
                "x-isnullable": false,
                "x-omitempty": false
              },
              "githubid": {
                "type": "string",
                "x-isnullable": false
              },
              "name": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        }
      }
    },
    "teams": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/team"
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        }
      }
    },
    "userDetails": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        },
        "teams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/team"
          }
        }
      }
    },
    "users": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/user"
      }
    }
  },
  "tags": [
    {
      "description": "Check if goliac is healthy",
      "name": "health"
    }
  ],
  "x-tagGroups": [
    {
      "name": "goliac Management",
      "tags": [
        "app"
      ]
    },
    {
      "name": "Health Check",
      "tags": [
        "health"
      ]
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "goliac is a golang REST application skeleton The base path for all the APIs is \"/api/v1\".\n",
    "title": "goliac",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/collaborators": {
      "get": {
        "description": "Get all external collaborators",
        "tags": [
          "app"
        ],
        "operationId": "getCollaborators",
        "responses": {
          "200": {
            "description": "get list of collarators",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/collaborators/{collaboratorID}": {
      "get": {
        "description": "Get collaborator and repos",
        "tags": [
          "app"
        ],
        "operationId": "getCollaborator",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "collaborator name",
            "name": "collaboratorID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get collaborator details especially repositories",
            "schema": {
              "$ref": "#/definitions/collaboratorDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flushcache": {
      "post": {
        "description": "Flush the Github remote cache",
        "tags": [
          "app"
        ],
        "operationId": "postFlushCache",
        "responses": {
          "200": {
            "description": "cache flushed"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/liveness": {
      "get": {
        "description": "Check if Goliac is healthy",
        "tags": [
          "health"
        ],
        "operationId": "getLiveness",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/readiness": {
      "get": {
        "description": "Check if Goliac is ready to serve",
        "tags": [
          "health"
        ],
        "operationId": "getReadiness",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories": {
      "get": {
        "description": "Get all repositories",
        "tags": [
          "app"
        ],
        "operationId": "getRepositories",
        "responses": {
          "200": {
            "description": "get list of repositories",
            "schema": {
              "$ref": "#/definitions/repositories"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{repositoryID}": {
      "get": {
        "description": "Get repository and associated teams",
        "tags": [
          "app"
        ],
        "operationId": "getRepository",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "repository slug name",
            "name": "repositoryID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get repository details especially teams that have access",
            "schema": {
              "$ref": "#/definitions/repositoryDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/resync": {
      "post": {
        "description": "Ask to sync again against Github",
        "tags": [
          "app"
        ],
        "operationId": "postResync",
        "responses": {
          "200": {
            "description": "resync in progress"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Get different statistics on Goliac",
        "tags": [
          "app"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "get Goliac statistics",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/teams": {
      "get": {
        "description": "Get all teams",
        "tags": [
          "app"
        ],
        "operationId": "getTeams",
        "responses": {
          "200": {
            "description": "get list of teams",
            "schema": {
              "$ref": "#/definitions/teams"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/teams/{teamID}": {
      "get": {
        "description": "Get team and associated users and repos",
        "tags": [
          "app"
        ],
        "operationId": "getTeam",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "team name",
            "name": "teamID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get team details especially users and repositories",
            "schema": {
              "$ref": "#/definitions/teamDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "Get all users",
        "tags": [
          "app"
        ],
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "get list of users (organization or external)",
            "schema": {
              "$ref": "#/definitions/users"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{userID}": {
      "get": {
        "description": "Get user and associated teams and repos",
        "tags": [
          "app"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "user name",
            "name": "userID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get user details especially teams and repositories",
            "schema": {
              "$ref": "#/definitions/userDetails"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "RepositoryDetailsCollaboratorsItems0": {
      "type": "object",
      "properties": {
        "access": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "RepositoryDetailsTeamsItems0": {
      "type": "object",
      "properties": {
        "access": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "TeamDetailsMembersItems0": {
      "type": "object",
      "properties": {
        "external": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "TeamDetailsOwnersItems0": {
      "type": "object",
      "properties": {
        "external": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "collaboratorDetails": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "repositories": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/repository"
      }
    },
    "repository": {
      "type": "object",
      "properties": {
        "allowUpdateBranch": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "archived": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "autoMergeAllowed": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "deleteBranchOnMerge": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "public": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        }
      }
    },
    "repositoryDetails": {
      "type": "object",
      "properties": {
        "allowUpdateBranch": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "archived": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "autoMergeAllowed": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "collaborators": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RepositoryDetailsCollaboratorsItems0"
          }
        },
        "deleteBranchOnMerge": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "public": {
          "type": "boolean",
          "x-isnullable": false,
          "x-omitempty": false
        },
        "teams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RepositoryDetailsTeamsItems0"
          }
        }
      }
    },
    "status": {
      "type": "object",
      "properties": {
        "detailedErrors": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "detailedWarnings": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "lastSyncError": {
          "type": "string"
        },
        "lastSyncTime": {
          "type": "string",
          "minLength": 1
        },
        "nbRepos": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbTeams": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbUsers": {
          "type": "integer",
          "x-omitempty": false
        },
        "nbUsersExternal": {
          "type": "integer",
          "x-omitempty": false
        },
        "version": {
          "type": "string"
        }
      }
    },
    "team": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "owners": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    },
    "teamDetails": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TeamDetailsMembersItems0"
          }
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        },
        "owners": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TeamDetailsOwnersItems0"
          }
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        }
      }
    },
    "teams": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/team"
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "name": {
          "type": "string",
          "x-isnullable": false
        }
      }
    },
    "userDetails": {
      "type": "object",
      "properties": {
        "githubid": {
          "type": "string",
          "x-isnullable": false
        },
        "repositories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/repository"
          }
        },
        "teams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/team"
          }
        }
      }
    },
    "users": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/user"
      }
    }
  },
  "tags": [
    {
      "description": "Check if goliac is healthy",
      "name": "health"
    }
  ],
  "x-tagGroups": [
    {
      "name": "goliac Management",
      "tags": [
        "app"
      ]
    },
    {
      "name": "Health Check",
      "tags": [
        "health"
      ]
    }
  ]
}`))
}
