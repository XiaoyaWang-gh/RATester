package jsconfig

import (
	"fmt"
	"sync"
	"testing"
)

func TestAddSourceRoot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Builder{
		sourceRootsMu: sync.RWMutex{},
		sourceRoots:   make(map[string]bool),
	}

	root := "testRoot"

	b.AddSourceRoot(root)

	b.sourceRootsMu.RLock()
	_, ok := b.sourceRoots[root]
	b.sourceRootsMu.RUnlock()

	if !ok {
		t.Errorf("Expected root to be added to sourceRoots")
	}

	b.AddSourceRoot(root)

	b.sourceRootsMu.RLock()
	_, ok = b.sourceRoots[root]
	b.sourceRootsMu.RUnlock()

	if ok {
		t.Errorf("Expected root not to be added to sourceRoots because it already exists")
	}
}
