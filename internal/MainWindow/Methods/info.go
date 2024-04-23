package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Printinfo(id int){
	l := postgres.PrintTable()
	for _, row := range l{
		if row.ID == id{
		fmt.Printf("\nYour ID: %d", row.ID)
            fmt.Printf("\nYour Name: %s", row.Name)
            fmt.Printf("\nYour Email: %s", row.Email)
            fmt.Printf("\nYour Pocket: %d\n", row.Pocket)
		}
	}
}