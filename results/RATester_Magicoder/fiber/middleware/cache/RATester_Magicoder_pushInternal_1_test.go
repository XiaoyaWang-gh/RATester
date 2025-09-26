package cache

import (
	"fmt"
	"testing"
)

func TestpushInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{},
		indices: []int{},
		maxidx:  0,
	}

	entry := heapEntry{
		key:   "testKey",
		exp:   123456789,
		bytes: 100,
		idx:   0,
	}

	h.pushInternal(entry)

	if len(h.entries) != 1 {
		t.Errorf("Expected length of entries to be 1, but got %d", len(h.entries))
	}

	if len(h.indices) != 1 {
		t.Errorf("Expected length of indices to be 1, but got %d", len(h.indices))
	}

	if h.entries[0] != entry {
		t.Errorf("Expected entry to be %v, but got %v", entry, h.entries[0])
	}

	if h.indices[0] != 0 {
		t.Errorf("Expected index to be 0, but got %d", h.indices[0])
	}
}
