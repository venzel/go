package main

import (
	"fmt"
	"validations/order"
)

func main() {
	product, err := order.NewProduct("Leite", "tiago@gmail.com.br", 190.45)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(product)
}
