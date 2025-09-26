package cache

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := indexedHeap{}
	if h.Len() != 0 {
		t.Errorf("Expected 0, got %d", h.Len())
	}
}
