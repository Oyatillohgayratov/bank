package mainwindow

import (
	m "BANK/internal/MainWindow/Methods"
	"BANK/storage/postgres"
	"fmt"
)

func Mainwindow(id int) bool{
	var s int
	for {
		fmt.Printf("\n0.Back\n1.My info\n2.Edit\n3.deposit money\n4.vizdraf money\n5.Send money\n6.Delate accaunt\n")
		fmt.Scan(&s)
		switch s {
		case 0:
			 return true
		case 1:
			m.Printinfo(id)
		case 2:
			_ = m.Chooseedit(id)
		case 3:
			m.Givedeposite(id)
		case 4:
			m.Givevizdraf(id)
		case 5:
			m.Sendmoney(id)
		case 6:
			postgres.DelateAccauntPsql(id)
			return true
		}

	}
}
