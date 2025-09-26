package nomad

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBuildProviders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ProviderBuilder{
		Configuration: Configuration{
			DefaultRule: "Host:traefik",
		},
		Namespaces: []string{"namespace1", "namespace2"},
	}

	providers := p.BuildProviders()

	assert.Len(t, providers, 2)
	assert.Equal(t, "namespace1", providers[0].namespace)
	assert.Equal(t, "namespace2", providers[1].namespace)
}
