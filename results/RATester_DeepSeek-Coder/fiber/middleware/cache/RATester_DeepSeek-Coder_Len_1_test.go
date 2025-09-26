package cache

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
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
		maxidx:  3,
	}

	expectedLen := 3
	actualLen := h.Len()

	if actualLen != expectedLen {
		t.Errorf("Expected length of heap to be %d, but got %d", expectedLen, actualLen)
	}
}
