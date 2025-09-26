package tailscale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	err := p.Init()
	require.NoError(t, err)
	require.NotNil(t, p.dynConfigs)
}
