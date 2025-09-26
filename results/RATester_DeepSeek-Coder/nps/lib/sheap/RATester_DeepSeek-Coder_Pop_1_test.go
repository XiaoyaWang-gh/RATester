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

	expected := int64(1)
	result := heap.Pop(h)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	expectedLen := 2
	if len(*h) != expectedLen {
		t.Errorf("Expected heap length %d, got %d", expectedLen, len(*h))
	}
}
