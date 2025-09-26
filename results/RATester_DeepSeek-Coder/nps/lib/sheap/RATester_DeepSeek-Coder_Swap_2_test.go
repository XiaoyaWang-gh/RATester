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

	h := &IntHeap{1, 2, 3, 4, 5}
	i, j := 1, 3
	(*h).Swap(i, j)

	expected := 4
	if (*h)[1] != int64(expected) {
		t.Errorf("Expected %d at position %d, got %d", expected, i, (*h)[1])
	}
	expected = 2
	if (*h)[3] != int64(expected) {
		t.Errorf("Expected %d at position %d, got %d", expected, j, (*h)[3])
	}
}
