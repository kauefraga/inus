package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const connString string = "postgresql://docker:docker@localhost:5432/inus?sslmode=disable"

func Connect() *sql.DB {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
