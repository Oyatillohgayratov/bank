package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Sendmoney(myid int){
	fmt.Print("ID: ")
	var id int
	fmt.Scan(&id)
	fmt.Printf("\nHow much do you want to send: ")
	var sum int
	fmt.Scan(&sum)

	m := postgres.GetPocket(id)	
	if m >= sum {
        postgres.SendMoneyPsql(sum,myid,id)
    } else {
        fmt.Printf("\nNot enough money\n")
    }
}