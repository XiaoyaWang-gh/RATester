package plugin

import (
	"fmt"
	"testing"
)

func TestNewManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := NewManager()

	if manager == nil {
		t.Error("NewManager() should not return nil")
	}

	if len(manager.loginPlugins) != 0 {
		t.Error("NewManager() should initialize loginPlugins as an empty slice")
	}

	// Add more test cases as needed
}
