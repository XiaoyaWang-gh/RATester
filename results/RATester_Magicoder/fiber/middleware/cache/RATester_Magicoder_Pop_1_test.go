package cache

import (
	"fmt"
	"testing"
)

func TestPop_1(t *testing.T) {
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
		indices: []int{1, 2, 3},
		maxidx:  3,
	}

	expected := heapEntry{key: "key3", exp: 3, bytes: 3, idx: 3}
	actual := h.Pop()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	if len(h.entries) != 2 {
		t.Errorf("Expected length of entries to be 2, but got %d", len(h.entries))
	}

	if len(h.indices) != 2 {
		t.Errorf("Expected length of indices to be 2, but got %d", len(h.indices))
	}

	if h.maxidx != 2 {
		t.Errorf("Expected maxidx to be 2, but got %d", h.maxidx)
	}
}
