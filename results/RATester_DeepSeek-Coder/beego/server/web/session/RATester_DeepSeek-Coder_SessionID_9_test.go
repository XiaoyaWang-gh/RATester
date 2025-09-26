package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &CookieSessionStore{sid: "test_sid"}

	sessionID := store.SessionID(ctx)
	if sessionID != "test_sid" {
		t.Errorf("Expected session ID to be 'test_sid', got '%s'", sessionID)
	}
}
