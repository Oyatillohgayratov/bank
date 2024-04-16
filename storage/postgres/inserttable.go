package postgres

import ("log"
	)

func InsertTable() {
	db := Connection()
	defer db.Close()

	insertTableSQL := `
		INSERT INTO "user" (name, email) VALUES ($1, $2)
	`

	stmt, err := db.Prepare(insertTableSQL)
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close()

	name := "kim"
	email := "kim@asdsa"
	_, err = stmt.Exec(name, email)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}
}



