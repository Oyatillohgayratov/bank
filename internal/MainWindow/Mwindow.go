package mainwindow

import (
	firstwindow "BANK/internal/FirstWindow"
	"BANK/internal/MainWindow/Methods"
	"fmt"
)


func Mainwindow(id int){
	var s int
	for {
		fmt.Printf("\n0.Back\n1.My info\n2.Edit\n3.deposit monney\n4.vizdraf monney\n5.Send money\n6.Delate accaunt\n")
		fmt.Scan(&s)
		switch s {
		case 0:
			firstwindow.Firstwindow()
		case 1:
			methods.Printinfo(id)
		case 2:
			_ = methods.Chooseedit(id)
		case 3:
			methods.Givedeposite(id)
		case 4:
			methods.Givevizdraf(id)
		}

	}
}