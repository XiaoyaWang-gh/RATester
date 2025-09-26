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

	sid := "test_sid"
	pder.sessions[sid] = pder.list.PushBack(sid)

	err := pder.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy failed: %v", err)
	}

	if _, ok := pder.sessions[sid]; ok {
		t.Errorf("SessionDestroy failed: session still exists")
	}
}
