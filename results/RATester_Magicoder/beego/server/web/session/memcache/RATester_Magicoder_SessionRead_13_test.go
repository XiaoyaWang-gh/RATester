package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &MemProvider{
		maxlifetime: 3600,
		conninfo:    []string{"localhost:11211"},
		poolsize:    10,
		password:    "password",
	}

	// Test case 1: Session not found
	_, err := rp.SessionRead(ctx, "non-existent-session")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Session found
	_, err = rp.SessionRead(ctx, "existing-session")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 3: Error in memcache connection
	rp.conninfo = []string{"invalid-host:11211"}
	_, err = rp.SessionRead(ctx, "session")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
