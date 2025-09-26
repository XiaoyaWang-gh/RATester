package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}

	err := p.Init()
	require.Error(t, err)
	require.Equal(t, "non-empty endpoint is required", err.Error())
}
