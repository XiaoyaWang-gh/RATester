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
		maxlifetime: 10,
		savePath:    "/tmp",
	}

	// Create a session that will expire in 5 seconds
	sid := "test_session"
	store := &MemSessionStore{
		sid:          sid,
		timeAccessed: time.Now().Add(-6 * time.Second),
		value:        make(map[interface{}]interface{}),
		lock:         sync.RWMutex{},
	}
	pder.list.PushBack(store)
	pder.sessions[sid] = pder.list.Back()

	// Call SessionGC
	pder.SessionGC(ctx)

	// Check if the session was removed
	if _, ok := pder.sessions[sid]; ok {
		t.Error("Session was not removed")
	}
}
