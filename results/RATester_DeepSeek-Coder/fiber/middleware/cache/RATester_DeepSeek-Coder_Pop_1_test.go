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
			{key: "test1", exp: 1, bytes: 1, idx: 0},
			{key: "test2", exp: 2, bytes: 2, idx: 1},
			{key: "test3", exp: 3, bytes: 3, idx: 2},
		},
		indices: []int{0, 1, 2},
		maxidx:  2,
	}

	popped := h.Pop()

	if len(h.entries) != 2 {
		t.Errorf("Expected length of entries to be 2, got %d", len(h.entries))
	}

	if popped != h.entries[len(h.entries)-1] {
		t.Errorf("Expected popped element to be the last element in the entries slice, got %v", popped)
	}

	if h.indices[0] != 1 || h.indices[1] != 2 {
		t.Errorf("Expected indices to be updated correctly, got %v", h.indices)
	}

	if h.maxidx != 2 {
		t.Errorf("Expected maxidx to be unchanged, got %d", h.maxidx)
	}
}
