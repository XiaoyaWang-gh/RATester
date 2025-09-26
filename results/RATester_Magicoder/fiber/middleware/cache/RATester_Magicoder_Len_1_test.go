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
			{key: "key1", exp: 1, bytes: 1, idx: 1},
			{key: "key2", exp: 2, bytes: 2, idx: 2},
			{key: "key3", exp: 3, bytes: 3, idx: 3},
		},
	}

	expectedLen := len(h.entries)
	actualLen := h.Len()

	if actualLen != expectedLen {
		t.Errorf("Expected Len() to return %d, but got %d", expectedLen, actualLen)
	}
}
