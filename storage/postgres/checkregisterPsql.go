package postgres

import (
	"fmt"
)

func CheckName(name string) bool {
	db := Connection()
	defer db.Close()

	var n string
	row2 := db.QueryRow("SELECT name from users where name = $1", name)

	_ = row2.Scan(&n)
	if len(n) != 0 {
		fmt.Printf("\nthis name is already registered\n")
		return true
	}

	return false
}

func CheckEmail(email string) bool {
	db := Connection()
	defer db.Close()

	var e string
	row2 := db.QueryRow("SELECT email from users where email = $1", email)

	_ = row2.Scan(&e)
	if len(e) != 0 {
		fmt.Printf("\nthis email is already registered\n")
		return true
	}

	return false
}
