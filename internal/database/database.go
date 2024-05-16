package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := "postgresql://docker:docker@localhost:5432/inus?sslmode=disable"

	if os.Getenv("APP_ENV") == "production" {
		connString = fmt.Sprintf(
			"user=%s password=%s host=aws-0-sa-east-1.pooler.supabase.com port=5432 dbname=postgres",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		)
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
