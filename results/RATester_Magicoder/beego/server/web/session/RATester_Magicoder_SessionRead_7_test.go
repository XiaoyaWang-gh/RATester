package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_7(t *testing.T) {
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

	sid := "test_session_id"
	_, err := pder.SessionRead(ctx, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if _, ok := pder.sessions[sid]; !ok {
		t.Errorf("Session %s not found in sessions map", sid)
	}
}
