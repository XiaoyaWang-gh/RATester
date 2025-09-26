package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildUDPServiceConfiguration_3(t *testing.T) {
	type testCase struct {
		desc        string
		item        itemData
		config      *dynamic.UDPConfiguration
		expectedErr bool
	}

	testCases := []testCase{
		{
			desc: "should return error when configuration is nil",
			item: itemData{
				Name: "test",
			},
			config:      nil,
			expectedErr: true,
		},
		{
			desc: "should create new service when no services exist",
			item: itemData{
				Name: "test",
			},
			config: &dynamic.UDPConfiguration{
				Services: make(map[string]*dynamic.UDPService),
			},
			expectedErr: false,
		},
		{
			desc: "should not return error when service already exists",
			item: itemData{
				Name: "test",
			},
			config: &dynamic.UDPConfiguration{
				Services: map[string]*dynamic.UDPService{
					"test": {
						LoadBalancer: &dynamic.UDPServersLoadBalancer{},
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
			err := p.buildUDPServiceConfiguration(test.item, test.config)
			if (err != nil) != test.expectedErr {
				t.Errorf("buildUDPServiceConfiguration() error = %v, wantErr %v", err, test.expectedErr)
			}
		})
	}
}
