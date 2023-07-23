//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"injections/product"
	"injections/usecases"

	"github.com/google/wire"
)

var productRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.RepositoryInterface), new(*product.Repository)),
)

func NewCreateProductUseCase(db *sql.DB) *usecases.CreateProductUseCase {
	wire.Build(
		productRepositoryDependency,
		usecases.NewCreateProductUseCase,
	)

	return &usecases.CreateProductUseCase{}
}

func NewGetProductUseCase(db *sql.DB) *usecases.GetProductUseCase {
	wire.Build(
		productRepositoryDependency,
		usecases.NewGetProductUseCase,
	)

	return &usecases.GetProductUseCase{}
}
