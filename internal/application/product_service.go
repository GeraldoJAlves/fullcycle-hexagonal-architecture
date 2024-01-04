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

func (p ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct(name, price)
	if _, err := product.IsValid(); err != nil {
		return nil, err
	}

	return p.Persistence.Save(product)
}

func (p ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return nil, err
	}

	return p.Persistence.Save(product)
}

func (p ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return nil, err
	}

	return p.Persistence.Save(product)
}
