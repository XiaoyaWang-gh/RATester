package gateway

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestRegisterBackendFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.RegisterBackendFuncs("group", "kind", func(name, namespace string) (string, *dynamic.Service, error) {
		return "backend", &dynamic.Service{}, nil
	})
	assert.Equal(t, map[string]map[string]BuildBackendFunc{
		"group": {
			"kind": func(name, namespace string) (string, *dynamic.Service, error) {
				return "backend", &dynamic.Service{}, nil
			},
		},
	}, p.groupKindBackendFuncs)
}
