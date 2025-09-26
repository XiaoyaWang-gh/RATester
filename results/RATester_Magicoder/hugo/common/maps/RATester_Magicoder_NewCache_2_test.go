package maps

import (
	"fmt"
	"testing"
)

func TestNewCache_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := NewCache[string, int]()

	if cache == nil {
		t.Error("Expected cache to be initialized, but got nil")
	}

	if cache.m == nil {
		t.Error("Expected cache.m to be initialized, but got nil")
	}
}
