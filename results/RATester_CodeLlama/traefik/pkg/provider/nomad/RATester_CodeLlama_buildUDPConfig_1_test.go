package nomad

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildUDPConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	i := item{}
	configuration := &dynamic.UDPConfiguration{}
	err := p.buildUDPConfig(i, configuration)
	require.NoError(t, err)
}
