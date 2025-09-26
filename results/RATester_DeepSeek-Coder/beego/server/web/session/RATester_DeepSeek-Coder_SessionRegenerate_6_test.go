package session

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSessionRegenerate_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &MemProvider{
		lock:        sync.RWMutex{},
		sessions:    make(map[string]*list.Element),
		list:        list.New(),
		maxlifetime: 1800,
		savePath:    "/tmp",
	}

	oldsid := "oldsid"
	sid := "sid"

	// Test case 1: Regenerate a session that exists
	element := pder.list.PushFront(&MemSessionStore{sid: oldsid, timeAccessed: time.Now(), value: make(map[interface{}]interface{})})
	pder.sessions[oldsid] = element
	store, err := pder.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if store.SessionID(ctx) != sid {
		t.Errorf("Expected new session ID %s, got %s", sid, store.SessionID(ctx))
	}
	if _, ok := pder.sessions[oldsid]; ok {
		t.Errorf("Expected old session ID %s to be deleted", oldsid)
	}

	// Test case 2: Regenerate a session that does not exist
	_, err = pder.SessionRegenerate(ctx, "notexists", sid)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
