package application_test

import (
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"

	"github.com/stretchr/testify/require"
)

func TestProduct_Getters(t *testing.T) {
	p := application.NewProduct("Ball", 9.99)

	require.NotNil(t, p.GetID())
	require.Equal(t, "Ball", p.GetName())
	require.Equal(t, 9.99, p.GetPrice())
	require.Equal(t, application.DISABLED, p.GetStatus())
}

func TestProduct_Enabled(t *testing.T) {
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

func TestProduct_Disabled(t *testing.T) {
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

func TestProduct_IsValid(t *testing.T) {
	p := application.NewProduct("Ball", 1)

	isValid, err := p.IsValid()

	require.Nil(t, err)
	require.Equal(t, true, isValid)
	require.Equal(t, application.DISABLED, p.GetStatus())

	p.Status = ""

	isValid, err = p.IsValid()

	require.Nil(t, err)
	require.Equal(t, true, isValid)
	require.Equal(t, application.DISABLED, p.GetStatus())

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
