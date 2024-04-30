package postgres

import "fmt"


func CheckUser(id int, password string) bool{
	db := Connection()
	defer db.Close()

	var i int
	var p string

	row2 := db.QueryRow("SELECT id,name from users where id = $1 and password = $2", id, password)

	err := row2.Scan(&i,&p)
	if err != nil {
		fmt.Printf("\nWrong ID or password\nTRYING again\n")
		return true
	}

	return false
}
