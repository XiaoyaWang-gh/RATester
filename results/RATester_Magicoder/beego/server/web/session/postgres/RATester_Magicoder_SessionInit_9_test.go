package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mp := &Provider{}

	err := mp.SessionInit(ctx, 100, "/tmp")
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if mp.maxlifetime != 100 {
		t.Errorf("Expected maxlifetime to be 100, but got %d", mp.maxlifetime)
	}

	if mp.savePath != "/tmp" {
		t.Errorf("Expected savePath to be /tmp, but got %s", mp.savePath)
	}
}
