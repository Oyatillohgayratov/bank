package main

import (
	firstwindow "BANK/internal/FirstWindow"
	mainwindow "BANK/internal/MainWindow"
	// _ "github.com/lib/pq"
)

func main() {
	id := firstwindow.Firstwindow()
	mainwindow.Mainwindow(id)
}