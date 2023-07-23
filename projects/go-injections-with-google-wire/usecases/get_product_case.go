package usecases

import "injections/product"

type GetProductUseCase struct {
	Repository product.RepositoryInterface
}

func NewGetProductUseCase(repository product.RepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{Repository: repository}
}

func (u *GetProductUseCase) Execute(name string) error {
	return nil
}
