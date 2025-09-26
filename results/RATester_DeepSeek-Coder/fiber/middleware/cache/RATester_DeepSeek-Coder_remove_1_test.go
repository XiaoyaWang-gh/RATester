package cache

import (
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
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

	key, bytes := h.remove(1)

	if key != "key2" || bytes != 2 {
		t.Errorf("Expected key 'key2' and bytes 2, got %s and %d", key, bytes)
	}

	if h.entries[1].key != "key3" || h.entries[1].bytes != 3 {
		t.Errorf("Expected key 'key3' and bytes 3 at position 1, got %s and %d", h.entries[1].key, h.entries[1].bytes)
	}

	if h.indices[1] != 1 {
		t.Errorf("Expected index 1 at position 1, got %d", h.indices[1])
	}
}
