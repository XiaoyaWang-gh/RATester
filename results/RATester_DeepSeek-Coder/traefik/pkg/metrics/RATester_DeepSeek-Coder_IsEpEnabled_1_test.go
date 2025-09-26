package metrics

import (
	"fmt"
	"testing"
)

func TestIsEpEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := &standardRegistry{epEnabled: true}

	if !registry.IsEpEnabled() {
		t.Errorf("Expected IsEpEnabled to return true, got false")
	}

	registry.epEnabled = false

	if registry.IsEpEnabled() {
		t.Errorf("Expected IsEpEnabled to return false, got true")
	}
}
