package idempotency

import (
	"fmt"
	"sync"
	"testing"
)

func TestLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	l := &MemoryLock{
		keys: make(map[string]*sync.Mutex),
	}

	key := "test_key"

	err := l.Lock(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, ok := l.keys[key]; !ok {
		t.Errorf("Expected key to be in the map")
	}

	err = l.Unlock(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, ok := l.keys[key]; ok {
		t.Errorf("Expected key to be removed from the map")
	}
}
