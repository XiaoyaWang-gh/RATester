package sheap

import (
	"fmt"
	"testing"
)

func TestLen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := IntHeap{1, 2, 3}
	if h.Len() != 3 {
		t.Errorf("h.Len() = %d, want 3", h.Len())
	}
}
