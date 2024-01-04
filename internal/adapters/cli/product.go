package cli

import (
	"fmt"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"
)

func Run(service application.ProductService, action string, productId string, productName string, productPrice float64) (string, error) {
	var result string = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been create with the price %f",
			product.GetID(), product.GetName(), product.GetPrice())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", product.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
