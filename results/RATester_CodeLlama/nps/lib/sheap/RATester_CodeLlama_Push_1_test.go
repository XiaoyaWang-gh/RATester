package sheap

import (
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := IntHeap{1, 2, 3}
	h.Push(4)
	if len(h) != 4 {
		t.Errorf("len(h) = %d, want 4", len(h))
	}
	if h[3] != 4 {
		t.Errorf("h[3] = %d, want 4", h[3])
	}
}
