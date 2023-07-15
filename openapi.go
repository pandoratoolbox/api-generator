package main

import (
	"strings"

	"github.com/tidwall/pretty"
)

var mapModels map[string]Struct
var mapEndpoints map[string]Endpoint

type Endpoint struct {
	Path          string
	Method        string
	RequestModel  string
	ResponseModel string
	HandlerFunc   string
	RestrictAuth  bool
	RestrictAdmin bool
}

func GenerateOpenAPI(structs []Struct) string {
	var spec string

	//generate schema objects

	schemas := `"schemas": {
`

	for _, s := range structs {
		schemas += `

"{{table_name_upper}}": {`
		for _, c := range s.Columns {
			if c.IsForeignObject {
				if strings.Contains(c.Type, "[]") {
					//array of objects

					schemas += `
"{{column_name_snake}}": {
					    "type": "array",
					    "items": {
					        "$ref": "#/components/schemas/{{foreign_table_name_upper}}"
					    }
					}`

				} else {
					//object

					schemas += `
"{{column_name_snake}}": {
						"$ref": "#/components/schemas/{{foreign_table_name_upper}}"
					}
					`

				}

				continue
			}

			switch c.Type {
			case "time.Time":
				schemas += `
"{{column_name_snake}}": {
					"type": "string",
					"format": "date-time"
				}
				`
			case "string":
				schemas += `
"{{column_name_snake}}": {
					"type": "string"
				}
				`
			case "int64":
				schemas += `
"{{column_name_snake}}": {
					"type": "integer",
					"format": "int64"
				}
				`
			case "float64":
				schemas += `
"{{column_name_snake}}": {
					"type": "number",
					"format": "float"
				}`
			case "bool":
				schemas += `
"{{column_name_snake}}": {
					"type": "boolean"
				}`
			}
		}

		schemas += `
}
`

	}

	schemas += `
}`

	//generate endpoints using schema objects in response and request
	spec = string(pretty.Pretty([]byte(spec)))
	return spec
}

// func GenerateOpenApiEndpointsFromStruct(s Struct) string {
// 	var out string

// 	path := "/" + s.NameSnake

// 	endpoint_base := `
// 	"/{{struct_name_snake}}": {
// 	%s
// 		}`

// 	endpoint_id := `
// 		"/{{struct_name_snake}}/:id": {
// 		%s
// 			}`

// 	base_c := `"post": {
// 		"summary": "Create a new {{struct_name_snake}}",
// 		"operationId": "create{{struct_name_upper}}",
// 		"tags": [
// 			"{{struct_name_snake}}",
// 			"create",
// 			"id"
// 		],
// 		"requestBody": {
// 			"required": true,
// 			"content": {
// 				"application/json": {
// 					"schema": {
// 						"$ref": "#/components/schemas/{{struct_name_upper}}"
// 					}
// 				}
// 			}
// 		},
// 		"responses": {
// 			"200": {
// 				"description": "Successful creation",
// 				"content": {
// 					"application/json": {
// 						"schema": {
// 							"$ref": "#/components/schemas/{{struct_name_upper}}"
// 						}
// 					}
// 				}
// 			},
// 			"default": {
// 				"description": "unexpected error",
// 				"content": {
// 					"application/json": {
// 						"schema": {
// 							"$ref": "#/components/schemas/Error"
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}`

// 	base_r := `"get": {
// 		"summary": "Get all {{struct_name_snake}}",
// 		"operationId": "get{{struct_name_upper}}",
// 		"tags": [
// 			"{{struct_name_snake}}",
// 			"list",
// 			"get"
// 		],
// 		"responses": {
// 			"200": {
// 				"description": "List of {{struct_name_snake}}",
// 				"content": {
// 					"application/json": {
// 						"schema": {
// 							"type": "array",
// 							"items": {
// 								"$ref": "#/components/schemas/{{struct_name_upper}}"
// 								}
// 							}
// 						}
// 					}
// 				}
// 			},
// 			"default": {
// 				"description": "unexpected error",
// 				"content": {
// 					"application/json": {
// 						"schema": {
// 							"$ref": "#/components/schemas/Error"
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}`

// 	// base_u :=

// 	// base_d :=

// 	return out
// }

// func GenerateOpenApiEndpointsAuth() string {
// 	var out string
// 	return out
// }
