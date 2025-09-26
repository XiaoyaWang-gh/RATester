package cache

import (
	"fmt"
	"testing"
)

func TestNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxEntries := 10
	cache := New(maxEntries)

	if cache.MaxEntries != maxEntries {
		t.Errorf("Expected MaxEntries to be %d, but got %d", maxEntries, cache.MaxEntries)
	}

	if cache.ll == nil {
		t.Error("Expected ll to not be nil")
	}

	if cache.OnEvicted != nil {
		t.Error("Expected OnEvicted to be nil")
	}
}
