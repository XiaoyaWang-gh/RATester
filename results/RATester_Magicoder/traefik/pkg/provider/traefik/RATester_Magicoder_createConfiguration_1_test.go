package traefik

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestCreateConfiguration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		staticCfg: static.Configuration{},
	}

	cfg := provider.createConfiguration(ctx)

	if cfg == nil {
		t.Error("Expected configuration, got nil")
	}

	if _, ok := cfg.HTTP.Services["noop"]; !ok {
		t.Error("Expected service 'noop' to be created")
	}
}
