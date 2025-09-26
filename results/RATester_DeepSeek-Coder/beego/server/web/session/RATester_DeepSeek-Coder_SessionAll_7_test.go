package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionAll_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 100,
		savePath:    "/tmp",
	}

	ctx := context.Background()
	total := fp.SessionAll(ctx)
	if total != 0 {
		t.Errorf("Expected 0, got %d", total)
	}
}
