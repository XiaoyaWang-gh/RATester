package runtime

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestGetTCPRoutersByEntryPoints_1(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc        string
		entryPoints []string
		config      *Configuration
		expected    map[string]map[string]*TCPRouterInfo
	}{
		{
			desc:        "no routers",
			entryPoints: []string{"web"},
			config:      &Configuration{TCPRouters: map[string]*TCPRouterInfo{}},
			expected:    map[string]map[string]*TCPRouterInfo{},
		},
		{
			desc:        "one router",
			entryPoints: []string{"web"},
			config: &Configuration{TCPRouters: map[string]*TCPRouterInfo{
				"router1": {
					TCPRouter: &dynamic.TCPRouter{
						EntryPoints: []string{"web"},
					},
				},
			}},
			expected: map[string]map[string]*TCPRouterInfo{
				"web": {
					"router1": {
						TCPRouter: &dynamic.TCPRouter{
							EntryPoints: []string{"web"},
						},
					},
				},
			},
		},
		// Add more test cases as needed
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.config.GetTCPRoutersByEntryPoints(ctx, test.entryPoints)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected: %v, but got: %v", test.expected, result)
			}
		})
	}
}
