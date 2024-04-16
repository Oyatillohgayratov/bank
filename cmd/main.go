package main

import (
	"fmt"
	"BANK/storage/postgres"

	_ "github.com/lib/pq"
)



func main() {
	postgres.CreateTable()
	// postgres.InsertTable()
	l := postgres.PrintTable()
	fmt.Println(string(l[0][0]))

}