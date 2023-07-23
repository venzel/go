package product

import (
	"errors"
)

type Product struct {
	Name  string
	Price float64
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{Name: name, Price: price}

	isValid := product.isValid(product)

	if !isValid {
		return nil, errors.New("invalid product")
	}

	return product, nil
}

func (p *Product) isValid(product *Product) bool {
	return product.Name != ""
}
