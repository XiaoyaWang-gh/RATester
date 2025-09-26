package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mp := &Provider{}

	// Test case 1: Session does not exist
	exists, err := mp.SessionExist(ctx, "non-existent-session")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if exists {
		t.Error("Expected session to not exist, but it does")
	}

	// Test case 2: Session exists
	exists, err = mp.SessionExist(ctx, "existing-session")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exists {
		t.Error("Expected session to exist, but it does not")
	}
}
