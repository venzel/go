package main

import "injections/entity"

func main() {
	productRepository := entity.NewProductRepository(nil)

	productService := entity.NewProductService(productRepository)

	productService.Create("Tiago", 20)
}
