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
		t.Errorf("Expected Len to be 3, but got %d", h.Len())
	}
}
