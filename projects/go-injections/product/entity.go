package product

import (
	"errors"
)

type Product struct {
	Name  string
	Price float64
}

func (p *Product) NewProduct(name string, price float64) (*Product, error) {
	product := &Product{Name: name, Price: price}

	isValid := p.isValid(product)

	if !isValid {
		return nil, errors.New("invalid product")
	}

	return product, nil
}

func (u *Product) isValid(product *Product) bool {
	return product.Name != ""
}
