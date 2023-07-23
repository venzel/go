package usecases

import "injections/product"

type CreateProductUseCase struct {
	Repository product.RepositoryInterface
}

func NewCreateProductUseCase(repository product.RepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{Repository: repository}
}

func (u *CreateProductUseCase) Execute(name string, price float64) error {
	return nil
}
