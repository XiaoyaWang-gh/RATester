package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mp := &Provider{
		maxlifetime: 10,
		savePath:    "/tmp",
	}

	// Test case 1: Session does not exist
	oldsid := "old_session_id"
	sid := "new_session_id"
	_, err := mp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Session exists
	_, err = mp.SessionRead(ctx, oldsid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 3: Session data is empty
	_, err = mp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 4: Session data is not empty
	_, err = mp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
