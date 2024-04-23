package postgres

import (
    "log"
    _ "github.com/lib/pq"
)

func Vizdraf(n,id int){
	db := Connection() 
    defer db.Close()

    _, err := db.Exec("UPDATE users SET pocket = pocket - $1 WHERE id = $2", n, id)
    if err!= nil {
        log.Fatal("Error updating the database:", err)
    }
}