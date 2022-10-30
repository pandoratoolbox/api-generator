package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"unicode"
)

func Init() error {
	return nil
}

//generate structs from postgres tables + relationships
//generate database/graphql functions (import models, graphjin, sql, context)
//generate store functions with db + cache support (import models, graphql, cache, context)
//generate handler functions (import models, store, json, http, chi, context)
//generate routing (import handler, chi)

type ForeignKey struct {
	Table         string
	Column        string
	ForeignTable  string
	ForeignColumn string
}

type Column struct {
	Table            string `db:"table"`
	Name             string `db:"name"`
	Type             string `db:"type"`
	IsNullable       bool
	IsUnique         bool
	IsForeignKey     bool
	ForeignKeyTable  string
	ForeignKeyColumn string
	IsForeignObject  bool
}

type Struct struct {
	Name      string
	NameUpper string
	NameSnake string
	Table     string
	Columns   map[string]Column
}

type StructMap map[string]Struct

func ToSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}

func ToUpperCase(s string) string {
	rs := []rune(s)
	new := []rune{}
	next_upper := true

	for _, r := range rs {
		if next_upper {
			new = append(new, unicode.ToUpper(r))
			next_upper = false
			continue
		}
		if unicode.IsPunct(r) {
			next_upper = true
			continue
		}

		new = append(new, r)
	}

	return string(new)
}

func Generate() error {
	err := os.Mkdir("./"+APP_NAME, 0777)
	if err != nil {
		return err
	}

	err = os.Mkdir("./client", 0777)
	if err != nil {
		return err
	}

	err = os.Mkdir("./client/models", 0777)
	if err != nil {
		return err
	}

	structs, err = GenerateStructs(conn)
	if err != nil {
		return err
	}

	fmt.Println(structs)

	err = GenerateTsClient(structs)
	if err != nil {
		return err
	}

	// panic("")

	m, err := GenerateModels(structs)
	if err != nil {
		return err
	}

	m = gen_models + GenerateImport(PACKAGE_TIME) + m

	err = os.Mkdir("./"+APP_NAME+"/models", 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./"+APP_NAME+"/models/models.go", []byte(m), 0777)
	if err != nil {
		return err
	}

	pg := GeneratePostgresConnection()

	pg = gen_connections + GenerateImport(PACKAGE_ENV, PACKAGE_SQL, PACKAGE_PGX, PACKAGE_OS, PACKAGE_LOG, PACKAGE_FMT, PACKAGE_ERRORS) + pg

	err = os.Mkdir("./"+APP_NAME+"/connections", 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./"+APP_NAME+"/connections/postgres.go", []byte(pg), 0777)
	if err != nil {
		return err
	}

	var r string

	err = os.Mkdir("./"+APP_NAME+"/graphql", 0777)
	if err != nil {
		return err
	}

	err = os.Mkdir("./"+APP_NAME+"/store", 0777)
	if err != nil {
		return err
	}

	err = os.Mkdir("./"+APP_NAME+"/handlers", 0777)
	if err != nil {
		return err
	}

	h_core := gen_handlers + GenerateImport(PACKAGE_MODELS, PACKAGE_STORE, PACKAGE_CHI, PACKAGE_CHI_JWTAUTH, PACKAGE_HTTP, PACKAGE_JSON, PACKAGE_STRCONV, PACKAGE_TIME, PACKAGE_CONTEXT, PACKAGE_ERRORS) + GenerateHandlerCore() + GenerateAuthHandlers()
	err = ioutil.WriteFile("./"+APP_NAME+"/handlers/handlers.go", []byte(h_core), 0777)
	if err != nil {
		return err
	}

	h_middleware := gen_handlers + GenerateImport(PACKAGE_MODELS, PACKAGE_HTTP, PACKAGE_JSON, PACKAGE_STRCONV, PACKAGE_TIME, PACKAGE_CONTEXT, PACKAGE_ERRORS) + GenerateCoreMiddleware()
	err = ioutil.WriteFile("./"+APP_NAME+"/handlers/middleware.go", []byte(h_middleware), 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./"+APP_NAME+"/graphql/init.go", []byte(graphql_init), 0777)
	if err != nil {
		return err
	}

	for _, s := range structs {

		g, err := GenerateGraphqlQueries(s)
		if err != nil {
			return err
		}

		g = gen_graphql + GenerateImport(PACKAGE_MODELS, PACKAGE_GRAPHJIN, PACKAGE_ERRORS, PACKAGE_SQL, PACKAGE_CONTEXT, PACKAGE_JSON) + g

		err = ioutil.WriteFile("./"+APP_NAME+"/graphql/"+ToSnakeCase(s.Name)+".go", []byte(g), 0777)
		if err != nil {
			return err
		}

		st, err := GenerateStoreFunctions(s, "graphql")
		if err != nil {
			return err
		}

		st = gen_store + GenerateImport(PACKAGE_MODELS, PACKAGE_GRAPHQL, PACKAGE_ERRORS, PACKAGE_CONTEXT) + st

		err = ioutil.WriteFile("./"+APP_NAME+"/store/"+ToSnakeCase(s.Name)+".go", []byte(st), 0777)
		if err != nil {
			return err
		}

		h, err := GenerateHandlers(s)
		if err != nil {
			return err
		}

		h = gen_handlers + GenerateImport(PACKAGE_MODELS, PACKAGE_STORE, PACKAGE_CHI, PACKAGE_CHI_JWTAUTH, PACKAGE_HTTP, PACKAGE_JSON, PACKAGE_STRCONV, PACKAGE_TIME, PACKAGE_CONTEXT, PACKAGE_ERRORS) + h

		err = ioutil.WriteFile("./"+APP_NAME+"/handlers/"+ToSnakeCase(s.Name)+".go", []byte(h), 0777)

		if err != nil {
			return err
		}
	}

	r = GenerateRoutes(structs)

	main := `package main

` + GenerateImport(PACKAGE_LOG, PACKAGE_CHI, PACKAGE_HANDLERS, PACKAGE_ERRORS, PACKAGE_CHI_CORS, PACKAGE_CHI_MIDDLEWARE) + `
func main() {` + r + `

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatalf("Error serving HTTP handlers: %v", err)
	}

}`

	fmt.Println(main)

	err = ioutil.WriteFile("./"+APP_NAME+"/main.go", []byte(main), 0777)

	if err != nil {
		return err
	}

	env, err := ioutil.ReadFile(".env")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./"+APP_NAME+"/.env", env, 0777)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", APP_NAME)
	cmd.Dir = "./" + APP_NAME

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	cmd = exec.Command("goimports", "-w", ".")
	cmd.Dir = "./" + APP_NAME
	//separate ModelData and Model structs, fix some errors
	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	cmd = exec.Command("go", "get")
	cmd.Dir = "./" + APP_NAME

	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	//get the correct version of custom json encoding library
	cmd = exec.Command("go", "get", "github.com/pandoratoolbox/json@21b1eb964277be3cdc85e9c761be521460e98260")
	cmd.Dir = "./" + APP_NAME

	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	cmd = exec.Command("go", "build", ".")
	cmd.Dir = "./" + APP_NAME

	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	return nil
}
