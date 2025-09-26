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

	config := newDynamicConfig()

	if config == nil {
		t.Error("Expected a non-nil dynamicConfig, but got nil")
	}

	if config.entryPoints == nil {
		t.Error("Expected a non-nil entryPoints map, but got nil")
	}

	if config.routers == nil {
		t.Error("Expected a non-nil routers map, but got nil")
	}

	if config.services == nil {
		t.Error("Expected a non-nil services map, but got nil")
	}
}
