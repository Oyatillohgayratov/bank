package postgres

import "log"

func GetPocket(id int) int {
	db := Connection()
	defer db.Close()
	var p int
	row2 := db.QueryRow("SELECT pocket from users where id = $1", id)

	err := row2.Scan(&p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
