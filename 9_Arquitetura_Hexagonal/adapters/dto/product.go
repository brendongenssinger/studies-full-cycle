package dto

import "github.com/codeedu/go-hexagonal/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if product == nil {
		return nil, nil
	}

	if p.ID != "" {
		product.ID = p.ID
	}
	if p.Name != "" {
		product.Name = p.Name
	}
	if p.Price != 0 {
		product.Price = p.Price
	}
	if p.Status != "" {
		product.Status = p.Status
	}

	_, err := product.IsValid()

	if err != nil {
		return &application.Product{}, err
	}

	return product, nil

}
