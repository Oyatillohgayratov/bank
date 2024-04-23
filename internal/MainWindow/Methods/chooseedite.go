package methods

import (
	"BANK/storage/postgres"
    "BANK/internal/hash"

	"fmt"
)

func Chooseedit(id int) error {
	var n int
	var s string
    for {
        fmt.Printf("\n0.Back\n1.Name\n2.Email\n3.Password\n")
        fmt.Scan(&n)
        switch n {
        case 0:
            return nil
        case 1:
			fmt.Printf("Enter new name: ")
			fmt.Scan(&s)
            postgres.Editname(s,id)
        case 2:
			fmt.Printf("Enter new email: ")
			fmt.Scan(&s)
            postgres.Editemail(s,id)
        case 3:
			fmt.Printf("Enter new password: ")
			fmt.Scan(&s)
            postgres.Editpassword(hash.HashPassword(s),id)
        }
    }
}	