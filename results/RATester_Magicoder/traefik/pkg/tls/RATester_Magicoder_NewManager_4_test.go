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

	manager := NewManager()

	if manager == nil {
		t.Error("Expected a non-nil manager, but got nil")
	}

	if manager.stores == nil {
		t.Error("Expected a non-nil stores map, but got nil")
	}

	if manager.configs == nil {
		t.Error("Expected a non-nil configs map, but got nil")
	}

	if len(manager.configs) != 1 {
		t.Errorf("Expected 1 config, but got %d", len(manager.configs))
	}

	if _, ok := manager.configs["default"]; !ok {
		t.Error("Expected a config with key 'default', but it was not found")
	}
}
