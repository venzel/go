package order

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Product struct {
	ID        string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"required,stringlength(5|255)"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
}

type Products []*Product

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		Name: name,
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
