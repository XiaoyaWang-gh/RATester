package sheap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 1)
	heap.Push(h, 2)
	heap.Push(h, 3)

	if h.Len() != 3 {
		t.Errorf("Expected length of heap to be 3, but got %d", h.Len())
	}

	if (*h)[0] != 1 {
		t.Errorf("Expected first element of heap to be 1, but got %d", (*h)[0])
	}

	if (*h)[1] != 2 {
		t.Errorf("Expected second element of heap to be 2, but got %d", (*h)[1])
	}

	if (*h)[2] != 3 {
		t.Errorf("Expected third element of heap to be 3, but got %d", (*h)[2])
	}
}
