package maps

import (
	"fmt"
	"testing"
)

func TestReset_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &SliceCache[int]{
		m: map[string][]int{
			"key1": {1, 2, 3},
			"key2": {4, 5, 6},
		},
	}

	cache.Reset()

	if len(cache.m) != 0 {
		t.Errorf("Expected cache to be empty after reset, but it's not")
	}
}
