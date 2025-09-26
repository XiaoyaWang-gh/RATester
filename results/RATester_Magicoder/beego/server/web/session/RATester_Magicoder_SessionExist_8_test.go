package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionExist_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 100,
		savePath:    "/tmp/sessions",
	}

	ctx := context.Background()

	// Test case 1: Session ID is empty
	exist, err := fp.SessionExist(ctx, "")
	if err == nil {
		t.Error("Expected error for empty session ID, but got nil")
	}
	if exist {
		t.Error("Expected session not to exist, but it does")
	}

	// Test case 2: Session ID is too short
	exist, err = fp.SessionExist(ctx, "a")
	if err == nil {
		t.Error("Expected error for too short session ID, but got nil")
	}
	if exist {
		t.Error("Expected session not to exist, but it does")
	}

	// Test case 3: Session ID exists
	exist, err = fp.SessionExist(ctx, "existing_session_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exist {
		t.Error("Expected session to exist, but it does not")
	}

	// Test case 4: Session ID does not exist
	exist, err = fp.SessionExist(ctx, "non_existing_session_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if exist {
		t.Error("Expected session not to exist, but it does")
	}
}
