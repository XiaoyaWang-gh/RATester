package sheap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := IntHeap{2, 1, 5}
	heap.Init(&h)

	if !h.Less(0, 1) {
		t.Errorf("Expected true, got false")
	}

	if h.Less(1, 0) {
		t.Errorf("Expected false, got true")
	}
}
