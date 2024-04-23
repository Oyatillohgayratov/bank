package mainwindow

import (
	f "BANK/internal/FirstWindow"
	"BANK/storage/postgres"
	m "BANK/internal/MainWindow/Methods"
	"fmt"
)


func Mainwindow(id int){
	var s int
	for {
		fmt.Printf("\n0.Back\n1.My info\n2.Edit\n3.deposit money\n4.vizdraf money\n5.Send money\n6.Delate accaunt\n")
		fmt.Scan(&s)
		switch s {
		case 0:
			f.Firstwindow()
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
			f.Firstwindow()
		}

	}
}