package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

var structs []Struct

var gen_models = `package models

`

// func GenerateContextKeys() string {
// 	var out string

// 	out = `

// type ctx_key int64

// const (
// 	CTX_is_auth = ctx_key(0)
// 	CTX_user_id = ctx_key(1)
// 	CTX_user_role = ctx_key(2)
// )`

// 	return out
// }

func GenerateStructs(db *sql.DB) ([]Struct, error) {

	res, err := db.QueryContext(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return structs, err
	}

	var tables []string
	for res.Next() {
		var table string
		err = res.Scan(&table)
		if err != nil {
			return structs, err
		}
		tables = append(tables, table)
	}

	struct_map := make(map[string]map[string]Column)

	for _, t := range tables {
		q := `SELECT column_name as name, udt_name::regtype as type 
FROM information_schema.columns
WHERE table_schema = 'public'
AND table_name = $1;`

		res, err = db.QueryContext(context.Background(), q, t)
		if err != nil {
			return structs, err
		}

		struct_map[t] = make(map[string]Column)

		for res.Next() {
			var c Column
			var tp string
			err = res.Scan(&c.Name, &c.Type)
			if err != nil {
				return structs, err
			}

			switch c.Type {
			case "jsonb":
				tp = "map[string]interface{}"
			case "double precision":
				tp = "float64"
			case "numeric":
				tp = "float64"
			case "integer":
				tp = "int64"
			case "bigint":
				tp = "int64"
			case "bigint[]":
				tp = "Ints"
			case "smallint":
				tp = "int64"
			case "smallint[]":
				tp = "Ints"
			case "character varying[]":
				tp = "Strings"
			case "text":
				tp = "string"
			case "text[]":
				tp = "Strings"
			case "double":
				tp = "float64"
			case "real":
				tp = "float64"
			case "character varying":
				tp = "string"
			case "timestamp without time zone":
				tp = "time.Time"
			case "timestamp with time zone":
				tp = "time.Time"
			case "boolean":
				tp = "bool"
			default:
				log.Fatal(c.Type)
			}

			c.Type = tp

			struct_map[t][c.Name] = c
			// if strings.Contains(c.Name, "_id") {
			// 	rt := ToUpperCase(strings.TrimSuffix(c.Name, "_id"))
			// 	struct_map[t][rt] = Column{
			// 		Name: rt,
			// 		Type: "[]" + rt,
			// 	}
			// }
		}

	}

	struct_map, err = GetUniqueKeys(struct_map)
	if err != nil {
		return structs, err
	}

	struct_map, err = GetForeignKeyRelationships(struct_map)
	if err != nil {
		return structs, err
	}

	for k, v := range struct_map {
		structs = append(structs, Struct{
			Name:    k,
			Columns: v,
		})
	}

	return structs, nil

}

func GenerateModels(structs []Struct) (string, error) {
	out := `package models`

	out = `
	
	type Strings []string
	type Ints []int64` + models_ctx

	for _, s := range structs {
		q := `
	
		type {{name}} struct {
{{name}}Data`

		qd := `

		type {{name}}Data struct {`

		for _, c := range s.Columns {
			if c.IsForeignObject {
				if strings.Contains(c.Type, "[]") {
					q += `
					` + ToUpperCase(c.Name) + "	" + "[]*" + strings.ReplaceAll(c.Type, "[]", "")
				} else {
					q += `
					` + ToUpperCase(c.Name) + "	*" + c.Type
				}

				continue
			}

			if strings.Contains(c.Type, "[]") {
				qd += `
				` + ToUpperCase(c.Name) + "	" + "[]*" + strings.ReplaceAll(c.Type, "[]", "")
			} else {
				qd += `
				` + ToUpperCase(c.Name) + "	*" + c.Type
			}

		}

		qd += `
		}`

		q += `
		}`

		qd = strings.ReplaceAll(qd, "{{name}}", ToUpperCase(s.Name))
		q = strings.ReplaceAll(q, "{{name}}", ToUpperCase(s.Name))

		out += qd + q
	}

	return out, nil
}

var models_ctx = `
type ctxkey int64

const (
	CTX_is_auth       = ctxkey(0)
	CTX_user_id       = ctxkey(1)
	CTX_user_role_ids      = ctxkey(2)
	CTX_user_timezone = ctxkey(3)
)
`
