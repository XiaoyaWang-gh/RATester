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

	t.Parallel()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 10,
		savePath:    "/tmp",
	}

	ctx := context.Background()

	fp.SessionGC(ctx)

	// Add assertions here to verify the behavior of the method.
}
