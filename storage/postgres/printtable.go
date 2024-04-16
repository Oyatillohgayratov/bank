package postgres

import (
	"fmt"
	"log"
)

func PrintTable() []string{
	db := Connection()
	defer db.Close()
	
	printTableSQL := `
		SELECT * FROM "user"
	`

	rows, err := db.Query(printTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	var l []string
	for rows.Next() {
		var id int
		var name string
		var email string

		rows.Scan(&id, &name, &email)
		
		n := fmt.Sprint(id,name,email)
		l = append(l,n)

		
	}
	return l
}
