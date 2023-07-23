package entity

import "database/sql"

type ProductRepositoryInterface interface {
	Create(product *Product) error
	List() []*Product
}

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) Create(product *Product) error {
	return nil
}

func (r *ProductRepository) List() []*Product {
	return nil
}
