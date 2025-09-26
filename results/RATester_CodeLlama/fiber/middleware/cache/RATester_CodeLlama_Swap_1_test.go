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

	h := indexedHeap{
		entries: []heapEntry{
			{key: "a", exp: 1, bytes: 1},
			{key: "b", exp: 2, bytes: 2},
			{key: "c", exp: 3, bytes: 3},
		},
		indices: []int{0, 1, 2},
		maxidx:  2,
	}
	h.Swap(0, 1)
	if h.entries[0].key != "b" || h.entries[1].key != "a" {
		t.Errorf("Swap failed")
	}
	if h.indices[0] != 1 || h.indices[1] != 0 {
		t.Errorf("Swap failed")
	}
}
