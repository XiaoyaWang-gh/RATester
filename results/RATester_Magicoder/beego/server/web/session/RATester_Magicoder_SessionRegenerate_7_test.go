package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		maxlifetime: 10,
		savePath:    "/tmp/sessions",
	}

	ctx := context.Background()

	// Test case 1: Regenerate session with new sid that does not exist
	oldsid := "old_sid"
	sid := "new_sid"
	_, err := fp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Regenerate session with new sid that already exists
	_, err = fp.SessionRegenerate(ctx, oldsid, sid)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Regenerate session with old sid that does not exist
	_, err = fp.SessionRegenerate(ctx, "nonexistent_sid", sid)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
