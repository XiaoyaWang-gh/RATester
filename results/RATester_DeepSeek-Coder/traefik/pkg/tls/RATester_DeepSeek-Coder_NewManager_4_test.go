package tls

import (
	"fmt"
	"testing"
)

func TestNewManager_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	manager := NewManager()

	if manager == nil {
		t.Fatal("Expected a non-nil manager, got nil")
	}

	if manager.stores == nil {
		t.Error("Expected a non-nil stores map, got nil")
	}

	if len(manager.stores) != 0 {
		t.Errorf("Expected an empty stores map, got %d elements", len(manager.stores))
	}

	if manager.configs == nil {
		t.Error("Expected a non-nil configs map, got nil")
	}

	if len(manager.configs) != 1 {
		t.Errorf("Expected a configs map with 1 element, got %d elements", len(manager.configs))
	}

	if _, ok := manager.configs["default"]; !ok {
		t.Error("Expected a configs map with a 'default' key")
	}
}
