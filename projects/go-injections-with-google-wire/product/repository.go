package product

import "database/sql"

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
	return nil
}

func (r *Repository) List() []*Product {
	return nil
}
