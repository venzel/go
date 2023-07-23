package main

import (
	"database/sql"
	"injections/product"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")

	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	usecase := product.NewProductUseCase(repository)

	usecase.Create("Tiago", 20)
}
