package mct

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(c string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", c)
	if err != nil {
		//trace
		return
	}
	//trace
	return
}
