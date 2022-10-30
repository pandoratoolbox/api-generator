package connections

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var Postgres *sql.DB

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

	Postgres, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Postgres error: %s", err)
	}

	retries := 5

	for r := 0; r < retries; r++ {
		err := Postgres.Ping()
		if err == nil {
			break
		}

		if r == retries {
			log.Fatalf("Unable to establish connection to Postgres: %s", err.Error())
		}

		time.Sleep(10 * time.Second)
	}
	return nil

}
