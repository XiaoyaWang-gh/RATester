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

	t.Parallel()

	cache := New(5)

	if cache.Len() != 0 {
		t.Errorf("Expected length of cache to be 0, got %d", cache.Len())
	}

	cache.Add("key1", "value1")

	if cache.Len() != 1 {
		t.Errorf("Expected length of cache to be 1, got %d", cache.Len())
	}

	cache.Add("key2", "value2")

	if cache.Len() != 2 {
		t.Errorf("Expected length of cache to be 2, got %d", cache.Len())
	}

	cache.Remove("key1")

	if cache.Len() != 1 {
		t.Errorf("Expected length of cache to be 1, got %d", cache.Len())
	}

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected length of cache to be 0, got %d", cache.Len())
	}
}
