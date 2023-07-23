package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

}
