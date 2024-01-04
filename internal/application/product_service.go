package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (p ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
