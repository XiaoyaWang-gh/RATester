package redis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	err := p.Init()
	require.Error(t, err)
	require.Equal(t, "unable to create client TLS configuration: ", err.Error())
}
