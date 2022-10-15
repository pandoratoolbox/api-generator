package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var conn *sql.DB

func InitPostgres() error {
	var err error
	if os.Getenv("POSTGRES_HOST") == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	db := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	q := "postgres://%s:%s@%s:%s/%s"
	connectionString := fmt.Sprintf(q, user, password, host, port, db)

	conn, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Postgres error: %s", err)
	}

	retries := 5

	for r := 0; r < retries; r++ {
		err := conn.Ping()
		if err == nil {
			break
		}

		if r == retries {
			log.Fatalf("Unable to establish connection to Postgres: %s", err.Error())
		}

		time.Sleep(10 * time.Second)
	}

	fmt.Println("Connected to postgres")
	return nil

}

func GetForeignKeyRelationships(struct_map map[string]map[string]Column) (map[string]map[string]Column, error) {
	var out map[string]map[string]Column

	q := `WITH unnested_confkey AS (
		SELECT oid, unnest(confkey) as confkey
		FROM pg_constraint
	  ),
	  unnested_conkey AS (
		SELECT oid, unnest(conkey) as conkey
		FROM pg_constraint
	  )
	  select
		tbl.relname                 AS constraint_table,
		col.attname                 AS constraint_column,
		referenced_tbl.relname      AS referenced_table,
		referenced_field.attname    AS referenced_column
	  FROM pg_constraint c
	  LEFT JOIN unnested_conkey con ON c.oid = con.oid
	  LEFT JOIN pg_class tbl ON tbl.oid = c.conrelid
	  LEFT JOIN pg_attribute col ON (col.attrelid = tbl.oid AND col.attnum = con.conkey)
	  LEFT JOIN pg_class referenced_tbl ON c.confrelid = referenced_tbl.oid
	  LEFT JOIN unnested_confkey conf ON c.oid = conf.oid
	  LEFT JOIN pg_attribute referenced_field ON (referenced_field.attrelid = c.confrelid AND referenced_field.attnum = conf.confkey)
	  WHERE c.contype = 'f';`

	res, err := conn.QueryContext(context.Background(), q)
	if err != nil {
		return out, err
	}

	for res.Next() {
		var fk ForeignKey

		err = res.Scan(&fk.Table, &fk.Column, &fk.ForeignTable, &fk.ForeignColumn)
		if err != nil {
			return out, err
		}

		new := Column{
			Table:            fk.Table,
			Name:             fk.Column,
			Type:             struct_map[fk.Table][fk.Column].Type,
			IsNullable:       struct_map[fk.Table][fk.Column].IsNullable,
			ForeignKeyTable:  fk.ForeignTable,
			ForeignKeyColumn: fk.ForeignColumn,
			IsForeignKey:     true,
			IsUnique:         struct_map[fk.Table][fk.Column].IsUnique,
		}

		struct_map[fk.Table][fk.Column] = new

		struct_map[fk.Table][fk.ForeignTable] = Column{
			Table:           fk.ForeignTable,
			Name:            ToUpperCase(strings.ReplaceAll(fk.Column, "_id", "")),
			Type:            ToUpperCase(fk.ForeignTable),
			IsForeignObject: true,
		}

		if !struct_map[fk.Table][fk.Column].IsUnique {
			struct_map[fk.ForeignTable][fk.Table] = Column{
				Table:           fk.Table,
				Name:            ToUpperCase(fk.Table),
				Type:            "[]" + ToUpperCase(fk.Table),
				IsForeignObject: true,
			}
		} else {
			// log.Fatal(fk.Table + " " + fk.Column)
			struct_map[fk.ForeignTable][fk.Table] = Column{
				Table:           fk.Table,
				Name:            ToUpperCase(fk.Table),
				Type:            ToUpperCase(fk.Table),
				IsForeignObject: true,
			}
		}

	}

	out = struct_map

	return out, nil

}

func GetUniqueKeys(struct_map map[string]map[string]Column) (map[string]map[string]Column, error) {
	var out map[string]map[string]Column
	q := `WITH unnested_confkey AS (
		SELECT oid, unnest(confkey) as confkey
		FROM pg_constraint
	  ),
	  unnested_conkey AS (
		SELECT oid, unnest(conkey) as conkey
		FROM pg_constraint
	  )
	  select
		tbl.relname                 AS constraint_table,
		col.attname                 AS constraint_column
	  FROM pg_constraint c
	  LEFT JOIN unnested_conkey con ON c.oid = con.oid
	  LEFT JOIN pg_class tbl ON tbl.oid = c.conrelid
	  LEFT JOIN pg_attribute col ON (col.attrelid = tbl.oid AND col.attnum = con.conkey)
	  LEFT JOIN pg_class referenced_tbl ON c.confrelid = referenced_tbl.oid
	  LEFT JOIN unnested_confkey conf ON c.oid = conf.oid
	  LEFT JOIN pg_attribute referenced_field ON (referenced_field.attrelid = c.confrelid AND referenced_field.attnum = conf.confkey)
	  WHERE c.contype = 'u';`

	res, err := conn.QueryContext(context.Background(), q)
	if err != nil {
		return out, err
	}

	for res.Next() {
		var t string
		var c string

		err = res.Scan(&t, &c)
		if err != nil {
			return out, err
		}

		struct_map[t][c] = Column{
			Type:             struct_map[t][c].Type,
			Name:             struct_map[t][c].Name,
			Table:            t,
			IsNullable:       struct_map[t][c].IsNullable,
			IsUnique:         true,
			IsForeignKey:     struct_map[t][c].IsForeignKey,
			ForeignKeyTable:  struct_map[t][c].ForeignKeyTable,
			ForeignKeyColumn: struct_map[t][c].ForeignKeyColumn,
		}
		if c == "id" {
			log.Fatal(struct_map[t][c])
		}

	}

	out = struct_map

	return out, nil
}
