package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildTCPServiceConfiguration_2(t *testing.T) {
	testCases := []struct {
		desc        string
		container   dockerData
		config      *dynamic.TCPConfiguration
		expectedErr bool
	}{
		{
			desc: "Happy path",
			container: dockerData{
				ID:          "123456",
				ServiceName: "service1",
				Health:      "healthy",
			},
			config: &dynamic.TCPConfiguration{
				Services: make(map[string]*dynamic.TCPService),
			},
			expectedErr: false,
		},
		{
			desc: "Container not healthy",
			container: dockerData{
				ID:          "123456",
				ServiceName: "service1",
				Health:      "unhealthy",
			},
			config: &dynamic.TCPConfiguration{
				Services: make(map[string]*dynamic.TCPService),
			},
			expectedErr: false,
		},
		{
			desc: "Nil configuration",
			container: dockerData{
				ID:          "123456",
				ServiceName: "service1",
				Health:      "healthy",
			},
			config:      nil,
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

			ctx := context.Background()
			builder := &DynConfBuilder{}
			err := builder.buildTCPServiceConfiguration(ctx, test.container, test.config)
			if (err != nil) != test.expectedErr {
				t.Errorf("buildTCPServiceConfiguration() error = %v, wantErr %v", err, test.expectedErr)
			}
		})
	}
}
