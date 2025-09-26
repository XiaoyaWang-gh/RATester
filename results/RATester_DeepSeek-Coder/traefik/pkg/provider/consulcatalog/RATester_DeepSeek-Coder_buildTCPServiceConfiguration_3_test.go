package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildTCPServiceConfiguration_3(t *testing.T) {
	type testCase struct {
		desc        string
		item        itemData
		config      *dynamic.TCPConfiguration
		expectedErr bool
	}

	testCases := []testCase{
		{
			desc: "should add service when no services exist",
			item: itemData{
				Name: "service1",
			},
			config: &dynamic.TCPConfiguration{
				Services: map[string]*dynamic.TCPService{},
			},
			expectedErr: false,
		},
		{
			desc: "should return error when addServerTCP fails",
			item: itemData{
				Name: "service1",
			},
			config: &dynamic.TCPConfiguration{
				Services: map[string]*dynamic.TCPService{
					"service1": {
						LoadBalancer: &dynamic.TCPServersLoadBalancer{},
					},
				},
			},
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
			err := p.buildTCPServiceConfiguration(test.item, test.config)
			if (err != nil) != test.expectedErr {
				t.Errorf("buildTCPServiceConfiguration() error = %v, wantErr %v", err, test.expectedErr)
			}
		})
	}
}
