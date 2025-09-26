package consul

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}

	err := p.Init()
	require.Error(t, err)
	require.Equal(t, "wildcard namespace is not supported", err.Error())
}
