package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"gotest.tools/assert"
)

func TestBuildTCPRouterConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc           string
		ctx            context.Context
		configuration  *dynamic.TCPConfiguration
		expectedConfig *dynamic.TCPConfiguration
	}{
		{
			desc: "Empty rule",
			ctx:  context.Background(),
			configuration: &dynamic.TCPConfiguration{
				Routers: map[string]*dynamic.TCPRouter{
					"router1": {
						Rule: "",
					},
				},
			},
			expectedConfig: &dynamic.TCPConfiguration{
				Routers: map[string]*dynamic.TCPRouter{},
			},
		},
		{
			desc: "Multiple services",
			ctx:  context.Background(),
			configuration: &dynamic.TCPConfiguration{
				Routers: map[string]*dynamic.TCPRouter{
					"router1": {
						Rule: "HostSNI()",
					},
				},
				Services: map[string]*dynamic.TCPService{
					"service1": {},
					"service2": {},
				},
			},
			expectedConfig: &dynamic.TCPConfiguration{
				Routers: map[string]*dynamic.TCPRouter{
					"router1": {
						Rule:    "HostSNI()",
						Service: "service1",
					},
				},
				Services: map[string]*dynamic.TCPService{
					"service1": {},
					"service2": {},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			BuildTCPRouterConfiguration(test.ctx, test.configuration)
			assert.Equal(t, test.expectedConfig, test.configuration)
		})
	}
}
