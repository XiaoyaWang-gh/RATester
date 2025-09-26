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

	l := &MemoryLock{
		keys: make(map[string]*sync.Mutex),
	}
	key := "key"
	l.Lock(key)
	l.Unlock(key)
	if _, ok := l.keys[key]; ok {
		t.Errorf("key should be deleted")
	}
}
