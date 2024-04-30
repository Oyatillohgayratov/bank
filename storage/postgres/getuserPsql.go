package postgres

import (
	"BANK/config"
	"log"
)

func GetUser(id int)config.User{
	var user config.User
	db := Connection() 
    defer db.Close()

	row := db.QueryRow("select name,email,pocket from users where id = $1",id)

	err := row.Scan(&user.Name,&user.Email,&user.Pocket)
	if err != nil{
		log.Fatal(err)
	}
	return user
}