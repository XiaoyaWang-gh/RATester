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
		t.Error("Expected a non-nil registry, but got nil")
	}

	if registry.IsEpEnabled() {
		t.Error("Expected IsEpEnabled to be false, but it was true")
	}

	if registry.IsRouterEnabled() {
		t.Error("Expected IsRouterEnabled to be false, but it was true")
	}

	if registry.IsSvcEnabled() {
		t.Error("Expected IsSvcEnabled to be false, but it was true")
	}

	// Add more tests for other methods in the Registry interface
}
