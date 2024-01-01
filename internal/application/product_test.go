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

func Test_ProductDisabled(t *testing.T) {
	p := application.Product{
		Name:   "Ball",
		Price:  0,
		Status: application.ENABLED,
	}

	err := p.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.Status)

	p.Status = application.ENABLED
	p.Price = 10

	err = p.Disable()
	require.Equal(t, application.ENABLED, p.Status)
	require.NotNil(t, err)
	require.Equal(t, "the prive must be zero in order to have the product disabled", err.Error())
}
