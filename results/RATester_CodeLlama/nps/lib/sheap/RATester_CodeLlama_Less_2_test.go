package sheap

import (
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := IntHeap{1, 2, 3, 4, 5}
	if !h.Less(0, 1) {
		t.Errorf("Less(0, 1) = false, want true")
	}
	if h.Less(1, 0) {
		t.Errorf("Less(1, 0) = true, want false")
	}
	if h.Less(0, 0) {
		t.Errorf("Less(0, 0) = true, want false")
	}
	if h.Less(1, 1) {
		t.Errorf("Less(1, 1) = true, want false")
	}
	if h.Less(2, 1) {
		t.Errorf("Less(2, 1) = true, want false")
	}
	if !h.Less(2, 3) {
		t.Errorf("Less(2, 3) = false, want true")
	}
	if !h.Less(3, 2) {
		t.Errorf("Less(3, 2) = false, want true")
	}
	if h.Less(3, 3) {
		t.Errorf("Less(3, 3) = true, want false")
	}
	if h.Less(4, 3) {
		t.Errorf("Less(4, 3) = true, want false")
	}
	if !h.Less(4, 5) {
		t.Errorf("Less(4, 5) = false, want true")
	}
	if !h.Less(5, 4) {
		t.Errorf("Less(5, 4) = false, want true")
	}
	if h.Less(5, 5) {
		t.Errorf("Less(5, 5) = true, want false")
	}
}
