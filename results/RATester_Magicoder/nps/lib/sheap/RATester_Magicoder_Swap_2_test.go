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

	h := IntHeap{1, 2, 3, 4, 5}
	i, j := 1, 3
	h.Swap(i, j)
	if h[i] != 4 || h[j] != 2 {
		t.Errorf("Swap failed. Expected: %v, %v. Got: %v, %v", 4, 2, h[i], h[j])
	}
}
