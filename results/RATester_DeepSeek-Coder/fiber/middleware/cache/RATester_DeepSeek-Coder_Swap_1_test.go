package cache

import (
	"fmt"
	"testing"
)

func TestSwap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{
			{key: "key1", exp: 1, bytes: 1, idx: 0},
			{key: "key2", exp: 2, bytes: 2, idx: 1},
		},
		indices: []int{0, 1},
		maxidx:  1,
	}

	h.Swap(0, 1)

	if h.entries[0].key != "key2" || h.entries[1].key != "key1" {
		t.Errorf("Swap failed. Expected: key2, key1. Got: %s, %s", h.entries[0].key, h.entries[1].key)
	}
	if h.indices[0] != 1 || h.indices[1] != 0 {
		t.Errorf("Index swap failed. Expected: 1, 0. Got: %d, %d", h.indices[0], h.indices[1])
	}
}
