package cache

import (
	"fmt"
	"testing"
)

func TestRemoveFirst_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{
			{key: "key1", exp: 1, bytes: 1, idx: 0},
			{key: "key2", exp: 2, bytes: 2, idx: 1},
			{key: "key3", exp: 3, bytes: 3, idx: 2},
		},
		indices: []int{0, 1, 2},
		maxidx:  2,
	}

	key, exp := h.removeFirst()

	if key != "key1" || exp != 1 {
		t.Errorf("Expected key 'key1' and expiration time 1, got %s and %d", key, exp)
	}

	if len(h.entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(h.entries))
	}

	if len(h.indices) != 2 {
		t.Errorf("Expected 2 indices, got %d", len(h.indices))
	}

	if h.maxidx != 1 {
		t.Errorf("Expected maxidx 1, got %d", h.maxidx)
	}
}
