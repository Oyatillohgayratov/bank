package postgres

import (
    "log"
)


func SendMoneyPsql(money, myID, id int) {
	db := Connection()
	defer db.Close()

	sendMoneySQL1 := `UPDATE users SET pocket = pocket - $1 WHERE id = $2;`
	_, err := db.Exec(sendMoneySQL1, money, myID)
	if err != nil {
		log.Fatal(err)
	}

	sendMoneySQL2 := `UPDATE users SET pocket = pocket + $1 WHERE id = $2;`
	_, err = db.Exec(sendMoneySQL2, money, id)
	if err != nil {
		log.Fatal(err)
	}
}
