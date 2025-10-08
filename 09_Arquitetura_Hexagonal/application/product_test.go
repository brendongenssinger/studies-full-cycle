package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.ID = "123"
	product.Name = "Product Test"
	product.Price = 10.0
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)

}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.ID = "123"
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	//require.Equal(t, "Product Test is disabled")

}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Price = 10.0
	product.Status = application.DISABLED
}
