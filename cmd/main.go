package main

import (
	firstwindow "BANK/internal/FirstWindow"
	mainwindow "BANK/internal/MainWindow"
)

func main() {
	for {
		id := firstwindow.Firstwindow()
		mainwindow.Mainwindow(id)
	}
}
