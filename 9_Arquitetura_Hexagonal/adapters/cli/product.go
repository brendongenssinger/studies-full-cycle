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
		result = fmt.Sprintf("Product %s created with id %s", product.Name, product.Id)
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s enabled", res.Name)

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s disabled", res.Name)
	default:
		rest, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s status: %s", rest.Name, rest.Status)

	}

	return result, nill

}
