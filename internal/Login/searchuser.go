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
	psw := hash.HashPassword(password)
	if postgres.CheckUser(id,psw){
		return SearchUser()
	}
	return id

	

	
}