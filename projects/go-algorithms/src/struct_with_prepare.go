package src

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        string
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(name string, price float64) *Order {
	order := &Order{
		Name:      name,
		Price:     price,
		UpdatedAt: time.Now(),
	}

	order.Prepare()

	return order
}

func (o *Order) Prepare() {
	o.ID = uuid.New().String()
	o.CreatedAt = time.Now()
}
