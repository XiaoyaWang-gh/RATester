package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	fp := &FileProvider{}

	err := fp.SessionInit(ctx, 10, "/tmp")
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if fp.maxlifetime != 10 {
		t.Errorf("Expected maxlifetime to be 10, but got %d", fp.maxlifetime)
	}

	if fp.savePath != "/tmp" {
		t.Errorf("Expected savePath to be /tmp, but got %s", fp.savePath)
	}
}
