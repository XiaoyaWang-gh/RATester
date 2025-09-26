package cache

import (
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: make([]heapEntry, 0),
		indices: make([]int, 0),
		maxidx:  0,
	}

	testEntry := heapEntry{
		key:   "testKey",
		exp:   123456789,
		bytes: 987654321,
		idx:   0,
	}

	h.Push(testEntry)

	if len(h.entries) != 1 {
		t.Errorf("Expected length of entries to be 1, got %d", len(h.entries))
	}

	if len(h.indices) != 1 {
		t.Errorf("Expected length of indices to be 1, got %d", len(h.indices))
	}

	if h.maxidx != 0 {
		t.Errorf("Expected maxidx to be 0, got %d", h.maxidx)
	}

	if h.entries[0] != testEntry {
		t.Errorf("Expected first entry to be %v, got %v", testEntry, h.entries[0])
	}

	if h.indices[0] != 0 {
		t.Errorf("Expected first index to be 0, got %d", h.indices[0])
	}
}
