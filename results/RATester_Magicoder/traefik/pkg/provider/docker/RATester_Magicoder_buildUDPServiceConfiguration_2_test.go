package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildUDPServiceConfiguration_2(t *testing.T) {
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
	configuration := &dynamic.UDPConfiguration{
		Services: make(map[string]*dynamic.UDPService),
	}

	err := (&DynConfBuilder{}).buildUDPServiceConfiguration(ctx, container, configuration)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(configuration.Services) != 1 {
		t.Errorf("Expected 1 service, but got: %v", len(configuration.Services))
	}
}
