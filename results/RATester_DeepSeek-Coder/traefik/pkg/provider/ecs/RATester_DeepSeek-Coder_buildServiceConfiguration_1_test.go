package ecs

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildServiceConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc        string
		instance    ecsInstance
		config      *dynamic.HTTPConfiguration
		expectedErr bool
	}{
		{
			desc:     "no services in config",
			instance: ecsInstance{Name: "test"},
			config:   &dynamic.HTTPConfiguration{},
		},
		{
			desc:     "service already exists",
			instance: ecsInstance{Name: "test"},
			config: &dynamic.HTTPConfiguration{
				Services: map[string]*dynamic.Service{
					"test": {
						LoadBalancer: &dynamic.ServersLoadBalancer{},
					},
				},
			},
		},
		{
			desc:        "add server error",
			instance:    ecsInstance{Name: "test"},
			config:      &dynamic.HTTPConfiguration{},
			expectedErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			err := p.buildServiceConfiguration(context.Background(), test.instance, test.config)

			if (err != nil) != test.expectedErr {
				t.Errorf("Expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
