package firstwindow

import (
	login "BANK/internal/Login"
	rg "BANK/internal/Registration"
	"BANK/storage/postgres"
	"fmt"
	"os"
)

func Firstwindow() int {
	n := 0
	for {
		fmt.Printf("\n1.Register\n2.Login\n3.Exite\n")
		fmt.Println("enter your choice: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			id := rg.Register()
			return id
		case 2:
			id := login.SearchUser()
			return id
		case 3:
			os.Exit(0)
		case 4:
			l := postgres.PrintTable()
			fmt.Println(l)
		default:
			fmt.Println("Invalid Input")
		}
	}
}
