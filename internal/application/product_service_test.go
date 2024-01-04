package application_test

import (
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"
	mock_application "github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	product := mock_application.NewMockProductInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	productService := application.NewProductService(persistence)

	productResult, err := productService.Get("uuid-1")

	require.Nil(t, err)
	require.Equal(t, product, productResult)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	product := mock_application.NewMockProductInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil)

	productService := application.NewProductService(persistence)

	productResult, err := productService.Create("ball", 1.99)

	require.Nil(t, err)
	require.Equal(t, product, productResult)
}
