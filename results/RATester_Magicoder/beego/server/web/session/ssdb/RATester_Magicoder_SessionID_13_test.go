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
	store := &SessionStore{sid: "test_sid"}

	sessionID := store.SessionID(ctx)

	if sessionID != "test_sid" {
		t.Errorf("Expected session ID to be 'test_sid', but got '%s'", sessionID)
	}
}
