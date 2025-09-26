package cache

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestRemoveInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{
			{"key1", 1, 1, 0},
			{"key2", 2, 2, 1},
			{"key3", 3, 3, 2},
		},
		indices: []int{0, 1, 2},
		maxidx:  2,
	}

	heap.Init(h)

	key, bytes := h.removeInternal(1)

	if key != "key2" || bytes != 2 {
		t.Errorf("Expected key 'key2' and bytes 2, got %s and %d", key, bytes)
	}

	if len(h.entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(h.entries))
	}

	if h.indices[0] != 0 || h.indices[1] != 2 {
		t.Errorf("Expected indices [0, 2], got %v", h.indices)
	}

	if h.maxidx != 2 {
		t.Errorf("Expected maxidx 2, got %d", h.maxidx)
	}
}
