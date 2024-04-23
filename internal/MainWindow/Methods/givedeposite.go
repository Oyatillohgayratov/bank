package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Givedeposite(id int){
	fmt.Printf("\nHow much do you want to deposit: ")
    var n int
    fmt.Scan(&n)
    postgres.Deposite(id,n)
}