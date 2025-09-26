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
			{key: "a", exp: 10, bytes: 100, idx: 0},
			{key: "b", exp: 20, bytes: 200, idx: 1},
			{key: "c", exp: 30, bytes: 300, idx: 2},
		},
	}

	if !h.Less(0, 1) {
		t.Error("Expected true, got false")
	}

	if h.Less(1, 0) {
		t.Error("Expected false, got true")
	}

	if h.Less(0, 2) {
		t.Error("Expected false, got true")
	}

	if !h.Less(2, 0) {
		t.Error("Expected true, got false")
	}
}
