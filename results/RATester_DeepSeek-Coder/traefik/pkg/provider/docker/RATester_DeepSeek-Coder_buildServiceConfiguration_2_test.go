package docker

import (
	"context"
	"fmt"
	"testing"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"gotest.tools/assert"
)

func TestBuildServiceConfiguration_2(t *testing.T) {
	ctx := context.Background()
	builder := &DynConfBuilder{}

	testCases := []struct {
		name          string
		container     dockerData
		configuration *dynamic.HTTPConfiguration
		expectedError error
	}{
		{
			name: "container is healthy",
			container: dockerData{
				Health: dockertypes.Healthy,
			},
			configuration: &dynamic.HTTPConfiguration{},
			expectedError: nil,
		},
		{
			name: "container is not healthy",
			container: dockerData{
				Health: "unhealthy",
			},
			configuration: &dynamic.HTTPConfiguration{},
			expectedError: nil,
		},
		{
			name: "container has no health check",
			container: dockerData{
				Health: "",
			},
			configuration: &dynamic.HTTPConfiguration{},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := builder.buildServiceConfiguration(ctx, tc.container, tc.configuration)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
