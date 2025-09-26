package cache

import (
	"fmt"
	"testing"
)

func Testremove_1(t *testing.T) {
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

	key, exp := h.remove(1)

	if key != "key2" || exp != 2 {
		t.Errorf("Expected key: 'key2', exp: 2, got key: '%s', exp: %d", key, exp)
	}

	if len(h.entries) != 2 || len(h.indices) != 2 {
		t.Errorf("Expected entries and indices length to be 2, got %d and %d", len(h.entries), len(h.indices))
	}

	if h.entries[0].key != "key1" || h.entries[1].key != "key3" {
		t.Errorf("Expected entries to be 'key1' and 'key3', got %s and %s", h.entries[0].key, h.entries[1].key)
	}

	if h.indices[0] != 0 || h.indices[1] != 2 {
		t.Errorf("Expected indices to be 0 and 2, got %d and %d", h.indices[0], h.indices[1])
	}
}
