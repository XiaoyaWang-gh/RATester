package ratelimit

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGetRemaining_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tokenBucket{
		RWMutex:     sync.RWMutex{},
		remaining:   10,
		capacity:    20,
		lastCheckAt: time.Now(),
		rate:        1 * time.Second,
	}

	remaining := b.getRemaining()
	if remaining != 10 {
		t.Errorf("Expected remaining to be 10, got %d", remaining)
	}
}
