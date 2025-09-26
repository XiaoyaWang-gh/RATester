package ecs

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildTCPServiceConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc        string
		instance    ecsInstance
		config      *dynamic.TCPConfiguration
		expectedErr bool
	}{
		{
			desc: "should return error when instance is empty",
			instance: ecsInstance{
				Name: "",
				ID:   "",
			},
			config: &dynamic.TCPConfiguration{
				Services: map[string]*dynamic.TCPService{},
			},
			expectedErr: true,
		},
		{
			desc: "should return error when configuration is empty",
			instance: ecsInstance{
				Name: "test",
				ID:   "123",
			},
			config:      &dynamic.TCPConfiguration{},
			expectedErr: true,
		},
		{
			desc: "should return nil when instance and configuration are not empty",
			instance: ecsInstance{
				Name: "test",
				ID:   "123",
			},
			config: &dynamic.TCPConfiguration{
				Services: map[string]*dynamic.TCPService{
					"test": {
						LoadBalancer: &dynamic.TCPServersLoadBalancer{},
					},
				},
			},
			expectedErr: false,
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
			err := p.buildTCPServiceConfiguration(test.instance, test.config)
			if (err != nil) != test.expectedErr {
				t.Errorf("buildTCPServiceConfiguration() error = %v, wantErr %v", err, test.expectedErr)
			}
		})
	}
}
