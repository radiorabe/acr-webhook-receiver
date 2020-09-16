// Code generated by go-swagger; DO NOT EDIT.

package acr

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
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hooks to receive ACRCloud events on and a simple API to query the received data.\nStores all data in a PostgreSQL database for later querying.\n",
    "title": "ACRCloud Webhooks Receiver",
    "license": {
      "name": "AGPLv3",
      "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/v1/_webhooks/results": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "This hook is called by the ACRCloud service when it knows what song we weere playing.",
        "tags": [
          "webhook"
        ],
        "summary": "ACRCloud results callback",
        "operationId": "addResult",
        "parameters": [
          {
            "description": "ACRCloud results entry",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Webhook"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/monitor-streams/{streamId}/results": {
      "get": {
        "description": "This endpoint implements the same API as upstream ACRCloud does.",
        "tags": [
          "compat"
        ],
        "summary": "ACRCloud Custom Streams Full Day Endpoint",
        "operationId": "getCustomStream",
        "parameters": [
          {
            "type": "string",
            "default": "s-qXuJARB",
            "description": "Stream ID, default is the non-realtime RaBe program.",
            "name": "streamId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Ignored but available for compatibility with upstream.",
            "name": "access_key",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The UTC date in the format 'YYYYMMDD'",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Results without local ID",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Data"
              }
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/results": {
      "get": {
        "description": "This is endpoint is useful for looking into and exporting the dataset.",
        "tags": [
          "api"
        ],
        "summary": "Get ACRCloud Results",
        "operationId": "getResults",
        "parameters": [
          {
            "$ref": "#/parameters/offsetParam"
          },
          {
            "$ref": "#/parameters/limitParam"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "from",
            "in": "query"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "to",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns array of results",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Result"
              }
            }
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/results/{resultId}": {
      "get": {
        "description": "Use this endpoint to fetch information on an exact entry.",
        "tags": [
          "api"
        ],
        "summary": "ACRCloud result",
        "operationId": "getResult",
        "parameters": [
          {
            "type": "integer",
            "name": "resultId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns single result",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Album": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "Artist": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "Data": {
      "type": "object",
      "required": [
        "status",
        "metadata"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/Metadata"
        },
        "result_type": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "$ref": "#/definitions/Status"
        }
      }
    },
    "ExternalIds": {
      "type": "object",
      "properties": {
        "deezer": {
          "type": "string"
        },
        "isrc": {
          "type": "string"
        },
        "itunes": {
          "type": "string"
        },
        "lyricfind": {
          "type": "string"
        },
        "musicstory": {
          "type": "string"
        },
        "spotify": {
          "type": "string"
        },
        "upc": {
          "type": "string"
        },
        "youtube": {
          "type": "string"
        }
      }
    },
    "ExternalMetadata": {
      "type": "object",
      "properties": {
        "deezer": {
          "type": "object"
        },
        "isrc": {
          "type": "object"
        },
        "itunes": {
          "type": "object"
        },
        "lyricfind": {
          "type": "object"
        },
        "musicstory": {
          "type": "object"
        },
        "spotify": {
          "type": "object"
        },
        "upc": {
          "type": "object"
        },
        "youtube": {
          "type": "object"
        }
      }
    },
    "Metadata": {
      "type": "object",
      "required": [
        "timestamp_utc",
        "played_duration"
      ],
      "properties": {
        "music": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Music"
          }
        },
        "played_duration": {
          "type": "integer",
          "format": "int64"
        },
        "timestamp_utc": {
          "type": "string"
        }
      }
    },
    "Music": {
      "type": "object",
      "required": [
        "acrid",
        "result_from",
        "score",
        "title",
        "play_offset_ms",
        "duration_ms",
        "sample_begin_time_offset_ms",
        "sample_end_time_offset_ms",
        "db_begin_time_offset_ms",
        "external_ids",
        "external_metadata"
      ],
      "properties": {
        "acrid": {
          "type": "string"
        },
        "album": {
          "$ref": "#/definitions/Album"
        },
        "artists": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Artist"
          }
        },
        "db_begin_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "duration_ms": {
          "type": "integer",
          "format": "int64"
        },
        "external_ids": {
          "$ref": "#/definitions/ExternalIds"
        },
        "external_metadata": {
          "$ref": "#/definitions/ExternalMetadata"
        },
        "label": {
          "type": "string"
        },
        "play_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "release_date": {
          "type": "string",
          "format": "date"
        },
        "result_from": {
          "type": "integer",
          "format": "int32"
        },
        "sample_begin_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "sample_end_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "score": {
          "type": "integer",
          "maximum": 100
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "x-go-custom-tag": "gorm:\"primaryKey;\""
        },
        "result": {
          "x-go-custom-tag": "gorm:\"type:jsonb;\"",
          "$ref": "#/definitions/Webhook"
        }
      }
    },
    "Status": {
      "type": "object",
      "required": [
        "msg",
        "version",
        "code"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "msg": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "Webhook": {
      "type": "object",
      "required": [
        "stream_id",
        "stream_url",
        "status",
        "data"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/Data"
        },
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "stream_id": {
          "type": "string"
        },
        "stream_url": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "limitParam": {
      "maximum": 50,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "The numbers of items to return.",
      "name": "limit",
      "in": "query"
    },
    "offsetParam": {
      "type": "integer",
      "default": 0,
      "description": "The number of items to skip before starting to collect the result set.",
      "name": "offset",
      "in": "query"
    }
  },
  "securityDefinitions": {
    "api_key": {
      "description": "Used to secure the Webhook endpoints.",
      "type": "apiKey",
      "name": "api_key",
      "in": "query"
    }
  },
  "externalDocs": {
    "description": "Find out more about Radio Bern RaBe",
    "url": "https://rabe.ch"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hooks to receive ACRCloud events on and a simple API to query the received data.\nStores all data in a PostgreSQL database for later querying.\n",
    "title": "ACRCloud Webhooks Receiver",
    "license": {
      "name": "AGPLv3",
      "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/v1/_webhooks/results": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "This hook is called by the ACRCloud service when it knows what song we weere playing.",
        "tags": [
          "webhook"
        ],
        "summary": "ACRCloud results callback",
        "operationId": "addResult",
        "parameters": [
          {
            "description": "ACRCloud results entry",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Webhook"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/monitor-streams/{streamId}/results": {
      "get": {
        "description": "This endpoint implements the same API as upstream ACRCloud does.",
        "tags": [
          "compat"
        ],
        "summary": "ACRCloud Custom Streams Full Day Endpoint",
        "operationId": "getCustomStream",
        "parameters": [
          {
            "type": "string",
            "default": "s-qXuJARB",
            "description": "Stream ID, default is the non-realtime RaBe program.",
            "name": "streamId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Ignored but available for compatibility with upstream.",
            "name": "access_key",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The UTC date in the format 'YYYYMMDD'",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Results without local ID",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Data"
              }
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/results": {
      "get": {
        "description": "This is endpoint is useful for looking into and exporting the dataset.",
        "tags": [
          "api"
        ],
        "summary": "Get ACRCloud Results",
        "operationId": "getResults",
        "parameters": [
          {
            "minimum": 0,
            "type": "integer",
            "default": 0,
            "description": "The number of items to skip before starting to collect the result set.",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 50,
            "minimum": 1,
            "type": "integer",
            "default": 20,
            "description": "The numbers of items to return.",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "from",
            "in": "query"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "to",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns array of results",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Result"
              }
            }
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    },
    "/v1/results/{resultId}": {
      "get": {
        "description": "Use this endpoint to fetch information on an exact entry.",
        "tags": [
          "api"
        ],
        "summary": "ACRCloud result",
        "operationId": "getResult",
        "parameters": [
          {
            "type": "integer",
            "name": "resultId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns single result",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "500": {
            "description": "Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Album": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "Artist": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "Data": {
      "type": "object",
      "required": [
        "status",
        "metadata"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/Metadata"
        },
        "result_type": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "$ref": "#/definitions/Status"
        }
      }
    },
    "ExternalIds": {
      "type": "object",
      "properties": {
        "deezer": {
          "type": "string"
        },
        "isrc": {
          "type": "string"
        },
        "itunes": {
          "type": "string"
        },
        "lyricfind": {
          "type": "string"
        },
        "musicstory": {
          "type": "string"
        },
        "spotify": {
          "type": "string"
        },
        "upc": {
          "type": "string"
        },
        "youtube": {
          "type": "string"
        }
      }
    },
    "ExternalMetadata": {
      "type": "object",
      "properties": {
        "deezer": {
          "type": "object"
        },
        "isrc": {
          "type": "object"
        },
        "itunes": {
          "type": "object"
        },
        "lyricfind": {
          "type": "object"
        },
        "musicstory": {
          "type": "object"
        },
        "spotify": {
          "type": "object"
        },
        "upc": {
          "type": "object"
        },
        "youtube": {
          "type": "object"
        }
      }
    },
    "Metadata": {
      "type": "object",
      "required": [
        "timestamp_utc",
        "played_duration"
      ],
      "properties": {
        "music": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Music"
          }
        },
        "played_duration": {
          "type": "integer",
          "format": "int64"
        },
        "timestamp_utc": {
          "type": "string"
        }
      }
    },
    "Music": {
      "type": "object",
      "required": [
        "acrid",
        "result_from",
        "score",
        "title",
        "play_offset_ms",
        "duration_ms",
        "sample_begin_time_offset_ms",
        "sample_end_time_offset_ms",
        "db_begin_time_offset_ms",
        "external_ids",
        "external_metadata"
      ],
      "properties": {
        "acrid": {
          "type": "string"
        },
        "album": {
          "$ref": "#/definitions/Album"
        },
        "artists": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Artist"
          }
        },
        "db_begin_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "duration_ms": {
          "type": "integer",
          "format": "int64"
        },
        "external_ids": {
          "$ref": "#/definitions/ExternalIds"
        },
        "external_metadata": {
          "$ref": "#/definitions/ExternalMetadata"
        },
        "label": {
          "type": "string"
        },
        "play_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "release_date": {
          "type": "string",
          "format": "date"
        },
        "result_from": {
          "type": "integer",
          "format": "int32"
        },
        "sample_begin_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "sample_end_time_offset_ms": {
          "type": "integer",
          "format": "int64"
        },
        "score": {
          "type": "integer",
          "maximum": 100,
          "minimum": 0
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "x-go-custom-tag": "gorm:\"primaryKey;\""
        },
        "result": {
          "x-go-custom-tag": "gorm:\"type:jsonb;\"",
          "$ref": "#/definitions/Webhook"
        }
      }
    },
    "Status": {
      "type": "object",
      "required": [
        "msg",
        "version",
        "code"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "msg": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "Webhook": {
      "type": "object",
      "required": [
        "stream_id",
        "stream_url",
        "status",
        "data"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/Data"
        },
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "stream_id": {
          "type": "string"
        },
        "stream_url": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "limitParam": {
      "maximum": 50,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "description": "The numbers of items to return.",
      "name": "limit",
      "in": "query"
    },
    "offsetParam": {
      "minimum": 0,
      "type": "integer",
      "default": 0,
      "description": "The number of items to skip before starting to collect the result set.",
      "name": "offset",
      "in": "query"
    }
  },
  "securityDefinitions": {
    "api_key": {
      "description": "Used to secure the Webhook endpoints.",
      "type": "apiKey",
      "name": "api_key",
      "in": "query"
    }
  },
  "externalDocs": {
    "description": "Find out more about Radio Bern RaBe",
    "url": "https://rabe.ch"
  }
}`))
}
