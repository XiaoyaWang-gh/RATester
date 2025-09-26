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
	expected := 3
	actual := h.Len()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
