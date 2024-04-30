package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Givevizdraf(id int){
	m := postgres.GetPocket(id)
	fmt.Printf("\nHow much do you want to vizdraf: ")
    var n int
    fmt.Scan(&n)
	if m >= n {
    postgres.Vizdraf(n,id)
	} else {
        fmt.Printf("\nYou don't have enough money\n")
    }
}