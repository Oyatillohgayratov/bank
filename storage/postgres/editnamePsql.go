package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Editname(name string,id int){
	db := Connection() 
    defer db.Close()

    _, err := db.Exec("UPDATE users SET name = $1 WHERE id = $2", name, id)
    if err!= nil {
        log.Fatal("Error updating the database:", err)
    }
    fmt.Println("Name updated")
}