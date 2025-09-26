package ingress

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Provider{}
	err := p.Init()
	require.NoError(t, err)
}
