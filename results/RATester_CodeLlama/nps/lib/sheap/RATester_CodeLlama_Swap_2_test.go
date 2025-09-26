package sheap

import (
	"fmt"
	"testing"
)

func TestSwap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := IntHeap{1, 2, 3}
	h.Swap(0, 1)
	if h[0] != 2 || h[1] != 1 {
		t.Errorf("Swap failed. Got %v, want %v", h, []int64{2, 1, 3})
	}
}
