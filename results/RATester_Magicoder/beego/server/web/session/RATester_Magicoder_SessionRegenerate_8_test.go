package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &CookieProvider{}

	oldsid := "old_session_id"
	sid := "new_session_id"

	store, err := pder.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("SessionRegenerate failed: %v", err)
	}

	if store == nil {
		t.Error("SessionRegenerate returned nil store")
	}

	// Add more assertions as needed
}
