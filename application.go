package main

import (
	"github.com/just1689/mysql-conn-test/mct"
	"time"
)

func main() {

	//Setup Jaeger

	//Get environment variables
	// TODO:
	connectionString := ""

	db, err := mct.Connect(connectionString)
	if err != nil {
		time.Sleep(1 * time.Minute)
		panic(err)
	}

	count, err := mct.QueryDate(db)
	if err != nil {
		time.Sleep(1 * time.Minute)
		panic(err)
	}

	//TODO: trace

	time.Sleep(1 * time.Minute)

}
