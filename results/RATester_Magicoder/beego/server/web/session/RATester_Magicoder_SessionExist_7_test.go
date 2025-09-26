package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &MemProvider{
		sessions: make(map[string]*list.Element),
	}

	// Test case 1: Session does not exist
	sid := "non-existing-session"
	exist, err := pder.SessionExist(ctx, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if exist {
		t.Errorf("Session %s should not exist", sid)
	}

	// Test case 2: Session exists
	sid = "existing-session"
	pder.sessions[sid] = &list.Element{}
	exist, err = pder.SessionExist(ctx, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exist {
		t.Errorf("Session %s should exist", sid)
	}
}
