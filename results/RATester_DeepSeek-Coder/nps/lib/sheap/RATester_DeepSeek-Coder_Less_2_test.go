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

	h := IntHeap{1, 2, 3}

	if !h.Less(0, 1) {
		t.Errorf("Expected true, got false")
	}

	if h.Less(1, 0) {
		t.Errorf("Expected false, got true")
	}

	if h.Less(1, 1) {
		t.Errorf("Expected false, got true")
	}
}
