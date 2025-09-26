package traefik

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateConfiguration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{}

	cfg := provider.createConfiguration(ctx)

	if cfg == nil {
		t.Fatal("Expected configuration, got nil")
	}

	if cfg.HTTP == nil {
		t.Error("Expected HTTP configuration, got nil")
	}

	if cfg.TCP == nil {
		t.Error("Expected TCP configuration, got nil")
	}

	if cfg.TLS == nil {
		t.Error("Expected TLS configuration, got nil")
	}

	if len(cfg.HTTP.Routers) == 0 {
		t.Error("Expected HTTP Routers, got none")
	}

	if len(cfg.TCP.Routers) == 0 {
		t.Error("Expected TCP Routers, got none")
	}

	if len(cfg.HTTP.Services) == 0 {
		t.Error("Expected HTTP Services, got none")
	}

	if len(cfg.TCP.Services) == 0 {
		t.Error("Expected TCP Services, got none")
	}
}
