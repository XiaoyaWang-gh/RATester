package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewMultiRegistry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registries := []Registry{
		&standardRegistry{
			epEnabled:     true,
			svcEnabled:    true,
			routerEnabled: true,
		},
		&standardRegistry{
			epEnabled:     true,
			svcEnabled:    true,
			routerEnabled: true,
		},
	}

	registry := NewMultiRegistry(registries)

	assert.True(t, registry.IsEpEnabled())
	assert.True(t, registry.IsSvcEnabled())
	assert.True(t, registry.IsRouterEnabled())
}
