package product

type ProductUseCaseInterface interface {
	Create(name string, price float64) (*Product, error)
	List() []*Product
}

type ProductUseCase struct {
	Repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{Repository: repository}
}

func (s *ProductUseCase) Create(name string, price float64) (*Product, error) {
	return nil, nil
}

func (s *ProductUseCase) List() []*Product {
	return nil
}
