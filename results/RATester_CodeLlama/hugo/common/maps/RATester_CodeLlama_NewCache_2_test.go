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

	cache := NewCache[string, string]()
	if cache == nil {
		t.Errorf("NewCache() returned nil")
	}
}
