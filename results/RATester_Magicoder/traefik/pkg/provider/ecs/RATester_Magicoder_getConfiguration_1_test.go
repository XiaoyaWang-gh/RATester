package ecs

import (
	"fmt"
	"testing"
)

func TestGetConfiguration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		ExposedByDefault: true,
	}

	instance := ecsInstance{
		Labels: map[string]string{
			"traefik.ecs.enable": "false",
		},
	}

	conf, err := provider.getConfiguration(instance)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if conf.Enable != false {
		t.Errorf("Expected Enable to be false, got %v", conf.Enable)
	}

	instance = ecsInstance{
		Labels: map[string]string{
			"traefik.ecs.enable": "true",
		},
	}

	conf, err = provider.getConfiguration(instance)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if conf.Enable != true {
		t.Errorf("Expected Enable to be true, got %v", conf.Enable)
	}

	instance = ecsInstance{
		Labels: map[string]string{
			"traefik.ecs.enable": "invalid",
		},
	}

	_, err = provider.getConfiguration(instance)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
