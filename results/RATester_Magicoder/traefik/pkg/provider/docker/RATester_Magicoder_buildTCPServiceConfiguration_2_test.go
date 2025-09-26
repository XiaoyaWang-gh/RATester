package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildTCPServiceConfiguration_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	container := dockerData{
		ID: "testID",
		Labels: map[string]string{
			"traefik.enable": "true",
		},
		Health: "healthy",
	}
	configuration := &dynamic.TCPConfiguration{
		Services: map[string]*dynamic.TCPService{},
	}

	err := (&DynConfBuilder{}).buildTCPServiceConfiguration(ctx, container, configuration)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(configuration.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(configuration.Services))
	}

	container.Health = "unhealthy"
	err = (&DynConfBuilder{}).buildTCPServiceConfiguration(ctx, container, configuration)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(configuration.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(configuration.Services))
	}
}
