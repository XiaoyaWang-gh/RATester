package acme

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	err := p.Init()
	require.Error(t, err)
	require.Equal(t, "unable to initialize ACME provider with no storage location for the certificates", err.Error())
}
