package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Editpassword(s string,id int){
	db := Connection() 
    defer db.Close()

    _, err := db.Exec("UPDATE users SET password = $1 WHERE id = $2", s, id)
    if err!= nil {
        log.Fatal("Error updating the database:", err)
    }
    fmt.Println("Password updated")
}