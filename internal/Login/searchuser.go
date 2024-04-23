package login

import (
	"BANK/storage/postgres"
	"BANK/internal/hash"
	"fmt"
)

func SearchUser() int{
	fmt.Printf("\nYour ID: ")
	var id int
	fmt.Scan(&id)
	fmt.Printf("Your Passward: ")
	var password string
	fmt.Scan(&password)

	l := postgres.PrintTable()
	for _, row := range l {
		if row.ID == id && row.Password == hash.HashPassword(password) {
            return id
        }
	}
	fmt.Printf("\nWrong ID or password\nTRYING again\n")
	return SearchUser()
	
}