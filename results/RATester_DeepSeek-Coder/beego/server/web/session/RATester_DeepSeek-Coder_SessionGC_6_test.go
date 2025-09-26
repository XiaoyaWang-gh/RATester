package session

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSessionGC_6(t *testing.T) {
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
		maxlifetime: 1,
		savePath:    "/tmp",
	}

	// Create a session
	sid := "test"
	store := &MemSessionStore{
		sid:          sid,
		timeAccessed: time.Now().Add(-2 * time.Second),
		value:        make(map[interface{}]interface{}),
		lock:         sync.RWMutex{},
	}
	pder.sessions[sid] = pder.list.PushBack(store)

	// Run GC
	pder.SessionGC(ctx)

	// Check if session was removed
	if _, ok := pder.sessions[sid]; ok {
		t.Errorf("Expected session to be removed, but it was not")
	}
}
