package session

import (
	"fmt"
	"testing"
)

func TestSessionID_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		config: &ManagerConfig{
			SessionIDLength: 16,
		},
	}

	sid, err := manager.sessionID()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(sid) != 32 {
		t.Errorf("Expected session ID length to be 32, got %d", len(sid))
	}

	manager.config.SessionIDLength = 0
	sid, err = manager.sessionID()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if len(sid) != 0 {
		t.Errorf("Expected empty session ID, got %s", sid)
	}
}
