package maps

import (
	"fmt"
	"testing"
)

func TestNewSliceCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := NewSliceCache[int]()

	if cache == nil {
		t.Error("Expected cache to be initialized, but it was nil")
	}

	if cache.m == nil {
		t.Error("Expected cache.m to be initialized, but it was nil")
	}
}
