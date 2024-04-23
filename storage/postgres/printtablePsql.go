package postgres

import (
	"BANK/config"
	"log"
)

func PrintTable() []config.User {
	db := Connection()
	defer db.Close()

	printTableSQL := `
		SELECT * FROM users
	`

	rows, err := db.Query(printTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []config.User
	for rows.Next() {
		var user config.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Pocket)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
