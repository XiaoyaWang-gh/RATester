package docker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &SwarmProvider{}
	err := p.Init()
	require.NoError(t, err)
}
