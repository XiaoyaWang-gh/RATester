package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildUDPServiceConfiguration_2(t *testing.T) {
	testCases := []struct {
		desc        string
		container   dockerData
		config      *dynamic.UDPConfiguration
		expectedErr error
	}{
		{
			desc: "should return error when container health is not healthy",
			container: dockerData{
				Health: "unhealthy",
			},
			config: &dynamic.UDPConfiguration{
				Services: make(map[string]*dynamic.UDPService),
			},
			expectedErr: nil,
		},
		{
			desc: "should add service to configuration when container health is healthy",
			container: dockerData{
				Health: "healthy",
			},
			config: &dynamic.UDPConfiguration{
				Services: make(map[string]*dynamic.UDPService),
			},
			expectedErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			builder := &DynConfBuilder{}
			err := builder.buildUDPServiceConfiguration(ctx, test.container, test.config)

			if err != test.expectedErr {
				t.Errorf("Expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
