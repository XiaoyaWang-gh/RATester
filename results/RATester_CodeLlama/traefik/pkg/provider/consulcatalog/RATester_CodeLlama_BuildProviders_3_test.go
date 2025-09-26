package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildProviders_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ProviderBuilder{
		Configuration: Configuration{
			Endpoint: &EndpointConfig{
				Address: "127.0.0.1:8500",
			},
		},
	}

	providers := p.BuildProviders()
	require.Len(t, providers, 1)
	require.Equal(t, "127.0.0.1:8500", providers[0].Endpoint.Address)
}
