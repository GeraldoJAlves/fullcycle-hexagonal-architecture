package application_test

import (
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"

	"github.com/stretchr/testify/require"
)

func Test_ProductEnabled(t *testing.T) {
	p := application.NewProduct("Ball", 9.99)

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
	p := application.NewProduct("Ball", 0)
	p.Status = application.ENABLED

	err := p.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.Status)

	p.Status = application.ENABLED
	p.Price = 10

	err = p.Disable()
	require.Equal(t, application.ENABLED, p.Status)
	require.NotNil(t, err)
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func Test_ProductIsValid(t *testing.T) {
	p := application.NewProduct("Ball", 1)

	isValid, err := p.IsValid()

	require.Nil(t, err)
	require.Equal(t, true, isValid)
	require.Equal(t, application.DISABLED, p.Status)

	p.Status = "wrong_status"

	isValid, err = p.IsValid()

	require.NotNil(t, err)
	require.Equal(t, "the status must be enabled or disabled", err.Error())
	require.Equal(t, false, isValid)

	p.Price = -1
	p.Status = application.DISABLED

	isValid, err = p.IsValid()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater or equal zero", err.Error())
	require.Equal(t, false, isValid)
}
