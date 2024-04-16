package postgres

import (
		"log"
	)

func CreateTable() {
	db := Connection()
	defer db.Close()

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT
		)
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}