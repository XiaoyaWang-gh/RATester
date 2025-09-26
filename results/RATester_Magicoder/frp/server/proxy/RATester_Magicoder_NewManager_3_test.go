package proxy

import (
	"fmt"
	"testing"
)

func TestNewManager_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := NewManager()

	if manager == nil {
		t.Error("NewManager() should not return nil")
	}

	if manager.pxys == nil {
		t.Error("NewManager() should initialize pxys")
	}
}
