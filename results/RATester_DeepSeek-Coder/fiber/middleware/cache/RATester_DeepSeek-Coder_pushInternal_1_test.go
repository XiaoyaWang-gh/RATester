package cache

import (
	"fmt"
	"testing"
)

func TestPushInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: make([]heapEntry, 0),
		indices: make([]int, 10),
		maxidx:  0,
	}

	entry := heapEntry{
		key:   "testKey",
		exp:   123456789,
		bytes: 987654321,
		idx:   0,
	}

	h.pushInternal(entry)

	if len(h.entries) != 1 {
		t.Errorf("Expected length of entries to be 1, got %d", len(h.entries))
	}

	if h.entries[0] != entry {
		t.Errorf("Expected entry to be %v, got %v", entry, h.entries[0])
	}

	if h.indices[0] != 0 {
		t.Errorf("Expected index of entry to be 0, got %d", h.indices[0])
	}
}
