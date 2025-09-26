package metrics

import (
	"fmt"
	"testing"
)

func TestNewVoidRegistry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := NewVoidRegistry()

	if registry == nil {
		t.Errorf("Expected a non-nil registry, got nil")
	}

	if registry.IsEpEnabled() {
		t.Errorf("Expected IsEpEnabled to be false, got true")
	}

	if registry.IsRouterEnabled() {
		t.Errorf("Expected IsRouterEnabled to be false, got true")
	}

	if registry.IsSvcEnabled() {
		t.Errorf("Expected IsSvcEnabled to be false, got true")
	}

	// Add more tests for other methods in the Registry interface
}
