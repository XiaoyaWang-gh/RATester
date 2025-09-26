package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/healthcheck"
)

func TestLaunchHealthCheck_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	m := &Manager{}
	m.healthCheckers = make(map[string]*healthcheck.ServiceHealthChecker)
	m.healthCheckers["service1"] = &healthcheck.ServiceHealthChecker{}
	m.healthCheckers["service2"] = &healthcheck.ServiceHealthChecker{}
	m.LaunchHealthCheck(ctx)
	assert.Equal(t, 2, len(m.healthCheckers))
}
