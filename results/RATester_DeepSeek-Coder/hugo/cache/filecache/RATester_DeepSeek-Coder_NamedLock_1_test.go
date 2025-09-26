package filecache

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/locker"
)

func TestNamedLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	cache := &Cache{
		nlocker: &lockTracker{
			seen:   make(map[string]struct{}),
			Locker: &locker.Locker{},
		},
	}

	id := "testID"
	unlock := cache.NamedLock(id)

	cache.nlocker.seenMu.RLock()
	_, ok := cache.nlocker.seen[id]
	cache.nlocker.seenMu.RUnlock()

	if !ok {
		t.Errorf("Expected id %s to be in seen map", id)
	}

	unlock()

	cache.nlocker.seenMu.RLock()
	_, ok = cache.nlocker.seen[id]
	cache.nlocker.seenMu.RUnlock()

	if ok {
		t.Errorf("Expected id %s to be removed from seen map", id)
	}
}
