package sheap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &IntHeap{1, 2, 3}
	heap.Init(h)
	expected := int64(3)
	actual := h.Pop()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
