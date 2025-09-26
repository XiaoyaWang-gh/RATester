package web

import (
	"fmt"
	"testing"
)

func TestNewHttpSever_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := NewHttpSever()
	if server == nil {
		t.Errorf("Expected non-nil server, got nil")
	}

	if server.Handlers == nil {
		t.Errorf("Expected non-nil Handlers, got nil")
	}

	if server.Server == nil {
		t.Errorf("Expected non-nil Server, got nil")
	}

	if server.Cfg == nil {
		t.Errorf("Expected non-nil Cfg, got nil")
	}

	if len(server.LifeCycleCallbacks) != 0 {
		t.Errorf("Expected empty LifeCycleCallbacks, got %d elements", len(server.LifeCycleCallbacks))
	}
}
