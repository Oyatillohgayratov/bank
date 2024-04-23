package methods

import (
	"BANK/storage/postgres"
	"fmt"
)

func Sendmoney(myid int){
	l := postgres.PrintTable()
	fmt.Print("ID: ")
	var id int
	fmt.Scan(&id)
	fmt.Printf("\nHow much do you want to send: ")
	var sum int
	fmt.Scan(&sum)

	m := 0
	for _, row := range l{
        if row.ID == myid{
            m = row.Pocket
        }
    }
	if m >= sum {
        postgres.SendMoneyPsql(sum,myid,id)
    } else {
        fmt.Printf("\nNot enough money\n")
    }
}