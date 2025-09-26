package metrics

import (
	"fmt"
	"testing"
)

func TestIsRouterEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := &standardRegistry{routerEnabled: true}

	if !registry.IsRouterEnabled() {
		t.Errorf("Expected IsRouterEnabled to return true, got false")
	}

	registry.routerEnabled = false

	if registry.IsRouterEnabled() {
		t.Errorf("Expected IsRouterEnabled to return false, got true")
	}
}
