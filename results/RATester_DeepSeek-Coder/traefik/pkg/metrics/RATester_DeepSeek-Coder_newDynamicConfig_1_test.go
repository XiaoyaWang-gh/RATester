package metrics

import (
	"fmt"
	"testing"
)

func TestNewDynamicConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := newDynamicConfig()

	if cfg == nil {
		t.Errorf("newDynamicConfig() = %v; want non-nil", cfg)
	}

	if cfg.entryPoints == nil {
		t.Errorf("newDynamicConfig().entryPoints = %v; want non-nil", cfg.entryPoints)
	}

	if cfg.routers == nil {
		t.Errorf("newDynamicConfig().routers = %v; want non-nil", cfg.routers)
	}

	if cfg.services == nil {
		t.Errorf("newDynamicConfig().services = %v; want non-nil", cfg.services)
	}
}
