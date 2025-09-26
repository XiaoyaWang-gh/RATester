package session

import (
	"fmt"
	"strings"
	"testing"
)

func TestSessionID_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{}
	manager.config = &ManagerConfig{}
	manager.config.SessionIDLength = 16
	manager.config.SessionIDPrefix = "sid"
	sid, err := manager.sessionID()
	if err != nil {
		t.Errorf("sessionID() failed: %v", err)
	}
	if len(sid) != 16 {
		t.Errorf("sessionID() failed: %v", sid)
	}
	if !strings.HasPrefix(sid, "sid") {
		t.Errorf("sessionID() failed: %v", sid)
	}
}
