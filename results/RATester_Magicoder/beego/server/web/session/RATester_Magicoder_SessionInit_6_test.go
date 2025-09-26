package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &MemProvider{}

	err := pder.SessionInit(ctx, 10, "/path")
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if pder.maxlifetime != 10 {
		t.Errorf("SessionInit failed: maxlifetime is not set correctly")
	}

	if pder.savePath != "/path" {
		t.Errorf("SessionInit failed: savePath is not set correctly")
	}
}
