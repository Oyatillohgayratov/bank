package postgres

import (
	"fmt"
	"log"
)

func DelateAccauntPsql(id int){
	db := Connection() 
    defer db.Close()

    _, err := db.Exec("DELETE FROM users WHERE id = $1", id)
    if err!= nil {
        log.Fatal("Error updating the database:", err)
    }else{
		fmt.Println("Accaunt deleted")
	}
}