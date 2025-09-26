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

	cache := NewSliceCache[string]()
	if cache == nil {
		t.Error("cache is nil")
	}
}
