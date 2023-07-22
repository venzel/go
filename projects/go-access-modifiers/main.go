package main

import "modifiers/orders"

func main() {
	product := orders.NewProduct("Keyboard", 100)
	price := product.GetPrice()
	println(price)
}
