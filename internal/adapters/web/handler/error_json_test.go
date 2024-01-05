package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandle_jsonError(t *testing.T) {
	b := jsonError("Teste error")

	require.Equal(t, []byte("{\"message\":\"Teste error\"}"), b)
}
