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

	l := &MemoryLock{
		keys: make(map[string]*sync.Mutex),
	}

	key := "testKey"

	err := l.Lock(key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	err = l.Unlock(key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := l.keys[key]
	if ok {
		t.Errorf("Expected key to be deleted, but it still exists")
	}
}
