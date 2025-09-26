package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cs := &SessionStore{
		sid: "test_session_id",
	}

	sessionID := cs.SessionID(ctx)

	if sessionID != "test_session_id" {
		t.Errorf("Expected session ID to be 'test_session_id', but got '%s'", sessionID)
	}
}
