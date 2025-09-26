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

	m := NewManager()
	if m == nil {
		t.Error("NewManager() returned nil")
	}
	if m.stores == nil {
		t.Error("NewManager() returned a nil stores map")
	}
	if m.configs == nil {
		t.Error("NewManager() returned a nil configs map")
	}
}
