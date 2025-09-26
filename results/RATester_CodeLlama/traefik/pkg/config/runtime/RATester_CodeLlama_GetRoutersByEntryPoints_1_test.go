package runtime

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestGetRoutersByEntryPoints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	c := &Configuration{
		Routers: map[string]*RouterInfo{
			"foo": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web"},
				},
			},
			"bar": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web", "websecure"},
				},
			},
			"baz": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web", "websecure", "websecure2"},
				},
			},
		},
	}

	entryPoints := []string{"web", "websecure"}
	tls := true
	expected := map[string]map[string]*RouterInfo{
		"web": {
			"foo": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web"},
				},
			},
			"bar": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web", "websecure"},
				},
			},
		},
		"websecure": {
			"bar": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web", "websecure"},
				},
			},
			"baz": {
				Router: &dynamic.Router{
					EntryPoints: []string{"web", "websecure", "websecure2"},
				},
			},
		},
	}

	assert.Equal(t, expected, c.GetRoutersByEntryPoints(ctx, entryPoints, tls))
}
