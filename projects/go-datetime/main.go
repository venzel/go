package main

import (
	"datetime/order"
	"fmt"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "UTC")
}

func main() {
	product, _ := order.NewProduct("Picanha", 190.45)

	fmt.Println(product)

	now := time.Now()

	day, month, year := now.Day(), int(now.Month()), now.Year()

	fmt.Println(day, month, year)

	fmt.Println(now.Unix()) // Timestamps in seconds
	fmt.Println(now.Format("02-01-2006 15:04:05"))

	p1, _ := order.NewProduct("Leite", 11.45)
	p2, _ := order.NewProduct("Sabão", 3.45)
	p3, _ := order.NewProduct("Café", 6.45)

	products := order.Products{p1, p2, p3}

	dateCreated := time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC)

	for _, product := range products {
		if product.CreatedAt.After(dateCreated) {
			fmt.Println(product)
		}
	}
}
