package postgres

import (
	"database/sql"
	"log"
)

func Connection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:azamat@localhost/bank?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return db
}