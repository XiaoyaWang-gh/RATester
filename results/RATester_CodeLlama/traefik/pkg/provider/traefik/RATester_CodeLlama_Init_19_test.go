package traefik

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &Provider{}
	err := i.Init()
	require.NoError(t, err)
}
