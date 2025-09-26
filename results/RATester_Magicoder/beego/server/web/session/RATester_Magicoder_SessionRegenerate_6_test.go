package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &MemProvider{
		sessions: make(map[string]*list.Element),
		list:     list.New(),
	}

	// Test case 1: Regenerate a session that exists
	oldsid := "old_session_id"
	sid := "new_session_id"
	element := pder.list.PushFront(&MemSessionStore{sid: oldsid, value: make(map[interface{}]interface{})})
	pder.sessions[oldsid] = element

	store, err := pder.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if store.SessionID(ctx) != sid {
		t.Errorf("Expected session ID %s, but got %s", sid, store.SessionID(ctx))
	}
	if _, ok := pder.sessions[oldsid]; ok {
		t.Errorf("Session with old ID %s should have been deleted", oldsid)
	}
	if _, ok := pder.sessions[sid]; !ok {
		t.Errorf("Session with new ID %s should exist", sid)
	}

	// Test case 2: Regenerate a session that does not exist
	_, err = pder.SessionRegenerate(ctx, "nonexistent_session_id", "new_session_id")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
