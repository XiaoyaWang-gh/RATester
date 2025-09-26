package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_6(t *testing.T) {
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

	// Add some sessions
	for i := 0; i < 10; i++ {
		sid := fmt.Sprintf("session_%d", i)
		pder.sessions[sid] = pder.list.PushBack(sid)
	}

	// Test SessionAll
	all := pder.SessionAll(ctx)
	if all != 10 {
		t.Errorf("Expected 10 sessions, got %d", all)
	}
}
