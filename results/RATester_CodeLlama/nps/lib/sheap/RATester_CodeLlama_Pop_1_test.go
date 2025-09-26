package sheap

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

	h := IntHeap{1, 2, 3}
	x := h.Pop()
	if x != 3 {
		t.Errorf("Pop() = %v, want 3", x)
	}
	if h.Len() != 2 {
		t.Errorf("Len() = %v, want 2", h.Len())
	}
	if h[0] != 1 || h[1] != 2 {
		t.Errorf("h[0] = %v, want 1; h[1] = %v, want 2", h[0], h[1])
	}
}
