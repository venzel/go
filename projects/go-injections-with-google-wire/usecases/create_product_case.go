package usecases

import (
	"fmt"
	"injections/product"
)

type CreateProductUseCase struct {
	Repository product.RepositoryInterface
}

func NewCreateProductUseCase(repository product.RepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{Repository: repository}
}

func (u *CreateProductUseCase) Execute(name string, price float64) error {
	fmt.Println("Executou case de criação de produtos no usecase ok!")

	u.Repository.Create(&product.Product{})

	return nil
}
