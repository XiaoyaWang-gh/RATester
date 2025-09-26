package cache

import (
	"fmt"
	"testing"
)

func TestNewMemoryCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := NewMemoryCache()
	if cache == nil {
		t.Error("NewMemoryCache() = nil")
	}
}
