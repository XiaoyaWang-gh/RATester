package idempotency

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnlock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	l := &MemoryLock{
		keys: make(map[string]*sync.Mutex),
	}

	key := "testKey"
	l.Lock(key)

	err := l.Unlock(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, ok := l.keys[key]
	if ok {
		t.Errorf("Expected key to be deleted, but it's still there")
	}
}
