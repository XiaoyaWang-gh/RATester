package metrics

import (
	"fmt"
	"testing"
)

func TestIsSvcEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := &standardRegistry{
		svcEnabled: true,
	}

	if !registry.IsSvcEnabled() {
		t.Errorf("Expected IsSvcEnabled to return true, got false")
	}

	registry.svcEnabled = false

	if registry.IsSvcEnabled() {
		t.Errorf("Expected IsSvcEnabled to return false, got true")
	}
}
