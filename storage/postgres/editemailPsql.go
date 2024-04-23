package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Editemail(email string,id int){
	db := Connection() 
    defer db.Close()

    _, err := db.Exec("UPDATE users SET email = $1 WHERE id = $2", email, id)
    if err!= nil {
        log.Fatal("Error updating the database:", err)
    }
    fmt.Println("Name updated")
}