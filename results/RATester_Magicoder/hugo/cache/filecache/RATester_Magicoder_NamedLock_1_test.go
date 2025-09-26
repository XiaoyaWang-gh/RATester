package filecache

import (
	"fmt"
	"testing"
)

func TestNamedLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		nlocker: &lockTracker{
			seen: make(map[string]struct{}),
		},
	}

	id := "testID"
	unlock := cache.NamedLock(id)

	// Test that the lock is acquired
	if _, ok := cache.nlocker.seen[id]; !ok {
		t.Errorf("Expected lock to be acquired for id %s", id)
	}

	// Test that the lock is released when the returned function is called
	unlock()
	if _, ok := cache.nlocker.seen[id]; ok {
		t.Errorf("Expected lock to be released for id %s", id)
	}
}
