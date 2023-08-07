package models

import (
	"database/sql"
	"fmt"
)

func ConnectDB() {
	fmt.Println("Entering here")
	dB, err := sql.Open("postgres", "host=localhost port=5432 user=test password=test dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = dB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}
