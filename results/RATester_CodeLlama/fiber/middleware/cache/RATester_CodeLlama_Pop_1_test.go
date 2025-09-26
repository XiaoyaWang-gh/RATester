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
			{key: "a", exp: 1, bytes: 1},
			{key: "b", exp: 2, bytes: 2},
			{key: "c", exp: 3, bytes: 3},
		},
		indices: []int{0, 1, 2},
		maxidx:  2,
	}
	if got := h.Pop(); got != "c" {
		t.Errorf("Pop() = %v, want %v", got, "c")
	}
	if got := h.Pop(); got != "b" {
		t.Errorf("Pop() = %v, want %v", got, "b")
	}
	if got := h.Pop(); got != "a" {
		t.Errorf("Pop() = %v, want %v", got, "a")
	}
}
