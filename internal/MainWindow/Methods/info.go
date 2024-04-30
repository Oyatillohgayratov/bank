package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Printinfo(id int){
	user := postgres.GetUser(id)
	fmt.Printf("\nYour ID: %d", id)
    fmt.Printf("\nYour Name: %s", user.Name)
    fmt.Printf("\nYour Email: %s", user.Email)
    fmt.Printf("\nYour Pocket: %d\n", user.Pocket)
}