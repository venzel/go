package orders

type product struct {
	Name  string
	price float64
}

func NewProduct(name string, price float64) *product {
	return &product{
		Name:  name,
		price: price,
	}
}

func (p product) GetPrice() float64 {
	return p.price
}
