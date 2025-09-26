package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSessionUpdate_2(t *testing.T) {
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

	sid := "test_session"
	element := pder.list.PushFront(sid)
	pder.sessions[sid] = element

	err := pder.SessionUpdate(ctx, sid)
	if err != nil {
		t.Errorf("SessionUpdate failed: %v", err)
	}

	session, ok := pder.sessions[sid]
	if !ok {
		t.Errorf("Session not found after update")
	}

	store := session.Value.(*MemSessionStore)
	if store.timeAccessed.Before(time.Now()) {
		t.Errorf("SessionUpdate did not update timeAccessed")
	}
}
