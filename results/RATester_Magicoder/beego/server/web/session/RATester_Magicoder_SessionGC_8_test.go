package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionGC_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 10,
		savePath:    "/tmp/sessions",
	}

	ctx := context.Background()
	fp.SessionGC(ctx)

	// Add assertions here
}
