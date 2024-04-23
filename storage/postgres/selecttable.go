package postgres

import (
	"BANK/internal/hash"
	"log"
)

func Select(name, email, password string) int {
	db := Connection()
	defer db.Close()
	hashpassword := hash.HashPassword(password)
	var id int

	row := db.QueryRow("SELECT id FROM users WHERE name = $1 AND email = $2 AND password = $3;", name, email, hashpassword)
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}
