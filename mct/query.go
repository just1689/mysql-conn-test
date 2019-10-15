package mct

import "database/sql"

func QueryDate(db *sql.DB) (count int, err error) {
	rows, err := db.Query("SELECT NOW()", nil)
	if err != nil {
		//trace
		return
	}
	count = 0
	for rows.Next() {
		count++
	}
	//trace
	return
}
