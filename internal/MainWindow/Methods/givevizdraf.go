package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Givevizdraf(id int){
	l := postgres.PrintTable()
	m := 0
	for _, row := range l{
		if row.ID == id{
            m = row.Pocket
        }
	}
	fmt.Printf("\nHow much do you want to vizdraf: ")
    var n int
    fmt.Scan(&n)
	if m >= n {
    postgres.Vizdraf(n,id)
	} else {
        fmt.Printf("\nYou don't have enough money\n")
    }
}