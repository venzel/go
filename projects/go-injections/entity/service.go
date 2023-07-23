package entity

type ProductServiceInterface interface {
	Create(name string, price float64) (*Product, error)
	List() []*Product
}

type ProductService struct {
	Repository ProductRepositoryInterface
}

func NewProductService(repository ProductRepositoryInterface) *ProductService {
	return &ProductService{Repository: repository}
}

func (s *ProductService) Create(name string, price float64) (*Product, error) {
	return nil, nil
}

func (s *ProductService) List() []*Product {
	return nil
}
