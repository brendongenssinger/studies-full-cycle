package cli

import (
	"fmt"

	"github.com/codeedu/go-hexagonal/adapters/cli"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productId, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productId).Return(productStatus, nil).AnyTimes()
	service.EXPECT().Disable(productId).Return(productStatus, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product Id %s & with the name", productId, productName)
	result, err := cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
