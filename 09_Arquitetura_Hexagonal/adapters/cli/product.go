package cli

import (
	"fmt"

	"github.com/codeedu/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s created with id %s", product.GetName(), product.GetID())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s enabled", res.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s disabled", res.GetName())

	default:
		rest, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s status: %s", rest.GetName(), rest.GetStatus())

	}

	return result, nil

}
