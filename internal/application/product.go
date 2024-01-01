package application

import "errors"

type ProductInterface interface {
	Enable()
}

type Product struct {
	Name   string
	Price  float64
	Status string
}

var (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}
