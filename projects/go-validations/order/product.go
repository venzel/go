package order

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Product struct {
	ID        string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"required,stringlength(5|255)"`
	Email     string    `json:"email" valid:"email"`
	Price     float64   `json:"price" valid:"required,range(0|100)~Preço acima do delimitado."`
	CreatedAt time.Time `json:"created_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct(name, email string, price float64) (*Product, error) {
	product := &Product{
		Name:  name,
		Email: email,
		Price: price,
	}

	product.prepare()

	err := product.isValid()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) prepare() {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}
