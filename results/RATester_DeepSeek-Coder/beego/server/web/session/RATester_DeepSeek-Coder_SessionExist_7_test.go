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

	t.Parallel()

	ctx := context.Background()
	pder := &MemProvider{
		sessions: make(map[string]*list.Element),
	}

	// Test with a non-existing session
	exist, err := pder.SessionExist(ctx, "non-existing-session")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if exist {
		t.Errorf("Expected session to not exist, but it does")
	}

	// Test with an existing session
	sid := "existing-session"
	pder.sessions[sid] = &list.Element{}
	exist, err = pder.SessionExist(ctx, sid)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !exist {
		t.Errorf("Expected session to exist, but it does not")
	}
}
