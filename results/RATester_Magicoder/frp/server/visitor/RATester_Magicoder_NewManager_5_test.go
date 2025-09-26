package visitor

import (
	"fmt"
	"testing"
)

func TestNewManager_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := NewManager()
	if manager == nil {
		t.Error("NewManager() should not return nil")
	}
	if manager.listeners == nil {
		t.Error("NewManager() should initialize listeners")
	}
}
