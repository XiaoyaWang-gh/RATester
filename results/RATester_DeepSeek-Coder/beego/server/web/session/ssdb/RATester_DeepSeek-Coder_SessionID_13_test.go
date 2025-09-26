package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		sid: "test_session_id",
	}

	sessionID := store.SessionID(ctx)
	if sessionID != "test_session_id" {
		t.Errorf("Expected session ID to be 'test_session_id', got '%s'", sessionID)
	}
}
