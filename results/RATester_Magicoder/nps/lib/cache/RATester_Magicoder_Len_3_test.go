package cache

import (
	"fmt"
	"testing"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := New(10)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	if cache.Len() != 3 {
		t.Errorf("Expected cache length to be 3, but got %d", cache.Len())
	}

	cache.Remove("key2")

	if cache.Len() != 2 {
		t.Errorf("Expected cache length to be 2 after removal, but got %d", cache.Len())
	}

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected cache length to be 0 after clear, but got %d", cache.Len())
	}
}
