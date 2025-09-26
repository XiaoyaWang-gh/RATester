package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/healthcheck"
)

func TestLaunchHealthCheck_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()

	manager := &Manager{
		healthCheckers: map[string]*healthcheck.ServiceHealthChecker{
			"service1": &healthcheck.ServiceHealthChecker{},
			"service2": &healthcheck.ServiceHealthChecker{},
		},
	}

	manager.LaunchHealthCheck(ctx)

	// Add assertions here to verify that the health checkers are launched
}
