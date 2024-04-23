package postgres

import "log"

func Deposite(id,n int){
	db := Connection()
    defer db.Close()

    plusbalanceSQL := `
        UPDATE users SET pocket = pocket + $1 WHERE id = $2
    `

    _, err := db.Exec(plusbalanceSQL, n, id)
    if err!= nil {
        log.Fatal(err)
    }
}