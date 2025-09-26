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

	h := &indexedHeap{}
	h.entries = []heapEntry{
		{key: "a", exp: 1, bytes: 1},
		{key: "b", exp: 2, bytes: 2},
		{key: "c", exp: 3, bytes: 3},
	}
	h.indices = []int{0, 1, 2}
	h.maxidx = 2
	a, b := h.removeFirst()
	if a != "a" || b != 1 {
		t.Errorf("removeFirst() = %v, %v; want %v, %v", a, b, "a", 1)
	}
	if h.maxidx != 2 {
		t.Errorf("removeFirst() = %v; want %v", h.maxidx, 2)
	}
	if h.indices[0] != 1 || h.indices[1] != 2 {
		t.Errorf("removeFirst() = %v; want %v", h.indices, []int{1, 2})
	}
	if h.entries[0].key != "b" || h.entries[1].key != "c" {
		t.Errorf("removeFirst() = %v; want %v", h.entries, []heapEntry{
			{key: "b", exp: 2, bytes: 2},
			{key: "c", exp: 3, bytes: 3},
		})
	}
}
