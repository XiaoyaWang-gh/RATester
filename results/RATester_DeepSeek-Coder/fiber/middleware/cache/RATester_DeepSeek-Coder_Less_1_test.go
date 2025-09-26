package cache

import (
	"fmt"
	"testing"
)

func TestLess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{
			{key: "a", exp: 1, bytes: 1, idx: 0},
			{key: "b", exp: 2, bytes: 2, idx: 1},
			{key: "c", exp: 3, bytes: 3, idx: 2},
		},
	}

	if !h.Less(0, 1) {
		t.Errorf("Expected Less(0, 1) to be true")
	}
	if h.Less(1, 0) {
		t.Errorf("Expected Less(1, 0) to be false")
	}
	if h.Less(1, 2) {
		t.Errorf("Expected Less(1, 2) to be false")
	}
	if !h.Less(2, 1) {
		t.Errorf("Expected Less(2, 1) to be true")
	}
}
