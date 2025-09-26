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

	if _, ok := b.sourceRoots[root]; !ok {
		t.Errorf("Expected root to be added, but it was not.")
	}

	b.AddSourceRoot(root)

	if len(b.sourceRoots) != 1 {
		t.Errorf("Expected root to be added only once, but it was added multiple times.")
	}
}
