package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var APP_NAME = "test"

func init() {
	//get app name from env APP_NAME = env.Get("APP_NAME")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	APP_NAME = os.Getenv("APP_NAME")
	PACKAGE_DATABASE = `"` + APP_NAME + "/database" + `"`
	PACKAGE_STORE = `"` + APP_NAME + "/store" + `"`
	PACKAGE_GRAPHQL = `"` + APP_NAME + "/graphql" + `"`
	PACKAGE_MODELS = `"` + APP_NAME + "/models" + `"`
	PACKAGE_HANDLERS = `"` + APP_NAME + "/handlers" + `"`
	PACKAGE_CONNECTIONS = `"` + APP_NAME + "/connections" + `"`
}

func main() {

	err := InitPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = Generate()
	if err != nil {
		log.Fatal(err)
	}
}
