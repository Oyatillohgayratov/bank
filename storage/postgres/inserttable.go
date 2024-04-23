package postgres

import (
	"BANK/internal/hash"
	"fmt"
	"log"
)

func InsertTable(name string, email string, password string) {
	db := Connection()
	defer db.Close()

	hashedPassword := hash.HashPassword(password) // Hashing the password

	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, hashedPassword)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}

	fmt.Println("Data inserted successfully.")

}
