package maps

import (
	"fmt"
	"testing"
)

func TestDrain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache[int, string]{
		m: map[int]string{
			1: "one",
			2: "two",
			3: "three",
		},
	}

	drained := cache.Drain()

	if len(drained) != 3 {
		t.Errorf("Expected 3 items, got %d", len(drained))
	}

	if len(cache.m) != 0 {
		t.Errorf("Expected cache to be empty, got %d items", len(cache.m))
	}
}
