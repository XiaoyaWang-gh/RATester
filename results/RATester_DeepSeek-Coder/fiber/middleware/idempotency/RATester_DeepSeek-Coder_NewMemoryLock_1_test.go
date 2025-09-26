package idempotency

import (
	"fmt"
	"testing"
)

func TestNewMemoryLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lock := NewMemoryLock()

	if lock == nil {
		t.Errorf("Expected NewMemoryLock to return a non-nil value")
	}

	if lock.keys == nil {
		t.Errorf("Expected keys to be initialized")
	}
}
