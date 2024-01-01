package application_test

import (
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"

	"github.com/stretchr/testify/require"
)

func Test_ProductEnabled(t *testing.T) {
	p := application.Product{
		Name:   "Ball",
		Price:  9.99,
		Status: application.DISABLED,
	}

	err := p.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, p.Status)

	p.Status = application.DISABLED
	p.Price = 0

	err = p.Enable()
	require.Equal(t, application.DISABLED, p.Status)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
