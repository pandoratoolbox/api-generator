package main

const (
	PACKAGE_ERRORS         = `"errors"`
	PACKAGE_FMT            = `"fmt"`
	PACKAGE_CONTEXT        = `"context"`
	PACKAGE_HTTP           = `"net/http"`
	PACKAGE_TIME           = `"time"`
	PACKAGE_STRCONV        = `"strconv"`
	PACKAGE_JSON           = `"github.com/pandoratoolbox/json"`
	PACKAGE_CHI            = `"github.com/go-chi/chi"`
	PACKAGE_CHI_CORS       = `"github.com/go-chi/cors"`
	PACKAGE_CHI_MIDDLEWARE = `"github.com/go-chi/chi/middleware"`
	PACKAGE_SQL            = `"database/sql"`
	PACKAGE_PGX            = `_ "github.com/jackc/pgx/v4/stdlib"`
	PACKAGE_OS             = `"os"`
	PACKAGE_REDIS          = `"github.com/go-redis/redis/v8"`
	PACKAGE_GRAPHJIN       = `"github.com/dosco/graphjin/core/v3"`
	PACKAGE_ENV            = `"github.com/joho/godotenv"`
	PACKAGE_LOG            = `"log"`
	PACKAGE_CHI_JWTAUTH    = `"github.com/go-chi/jwtauth"`
	PACKAGE_STRINGS        = `"strings"`
	PACKAGE_REFLECT        = `"reflect"`
	PACKAGE_UNICODE        = `"unicode"`
)

var (
	PACKAGE_MODELS      string
	PACKAGE_STORE       string
	PACKAGE_DATABASE    string
	PACKAGE_GRAPHQL     string
	PACKAGE_HANDLERS    string
	PACKAGE_CONNECTIONS string
)

func GetGoPackages() error {
	return nil
}

func GenerateImport(packages ...string) string {
	var out string
	out = `import (`

	for _, p := range packages {
		out += `
` + p
	}

	out += `
)

`

	return out
}
