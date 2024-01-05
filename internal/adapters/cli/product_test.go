package cli_test

import (
	"fmt"
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/adapters/cli"
	mock_application "github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Ball 2"
	productPrice := 2.99
	productStatus := "enabled"
	productId := "uuid-1"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(gomock.Any(), gomock.All()).Return(productMock, nil)

	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("Product ID %s with the name %s has been create with the price %f",
		productId, productName, productPrice), result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Ball 2"
	productPrice := 2.99
	productId := "uuid-1"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Get(gomock.Any()).Return(productMock, nil)
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil)

	result, err := cli.Run(serviceMock, "enable", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("Product %s has been enabled", productName), result)
}
