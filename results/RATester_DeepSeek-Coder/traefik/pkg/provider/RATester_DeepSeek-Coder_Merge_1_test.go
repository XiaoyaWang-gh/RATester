package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"gotest.tools/assert"
)

func TestMerge_1(t *testing.T) {
	tests := []struct {
		desc     string
		configs  map[string]*dynamic.Configuration
		expected *dynamic.Configuration
	}{
		{
			desc: "Given multiple configurations, when merging them, then the resulting configuration should be the one with the most specific settings",
			configs: map[string]*dynamic.Configuration{
				"config1": {
					HTTP: &dynamic.HTTPConfiguration{
						Routers: map[string]*dynamic.Router{
							"router1": {
								Rule: "Host(`example.com`)",
							},
						},
					},
				},
				"config2": {
					HTTP: &dynamic.HTTPConfiguration{
						Routers: map[string]*dynamic.Router{
							"router1": {
								Rule:    "Host(`example.com`)",
								Service: "service1",
							},
						},
					},
				},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"router1": {
							Rule:    "Host(`example.com`)",
							Service: "service1",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			result := Merge(ctx, test.configs)
			assert.Equal(t, test.expected, result)
		})
	}
}
