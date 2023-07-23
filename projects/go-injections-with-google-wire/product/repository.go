package product

import (
	"database/sql"
	"fmt"
)

type RepositoryInterface interface {
	Create(product *Product) error
	List() []*Product
}

type Repository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *Repository {
	return &Repository{Db: db}
}

func (r *Repository) Create(product *Product) error {
	fmt.Println("Executou case de criação de produtos no repositório ok!")

	return nil
}

func (r *Repository) List() []*Product {
	return nil
}
