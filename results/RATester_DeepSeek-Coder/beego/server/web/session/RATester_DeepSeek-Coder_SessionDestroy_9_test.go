package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_9(t *testing.T) {
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
	element := pder.list.PushBack(sid)
	pder.sessions[sid] = element

	err := pder.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, ok := pder.sessions[sid]; ok {
		t.Errorf("Expected session to be deleted, but it still exists")
	}

	if pder.list.Front() != nil {
		t.Errorf("Expected list to be empty, but it's not")
	}
}
