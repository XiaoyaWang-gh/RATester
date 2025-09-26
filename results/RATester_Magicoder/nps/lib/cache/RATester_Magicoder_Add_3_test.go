package cache

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := New(2)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	value, ok := cache.Get("key1")
	if ok {
		t.Errorf("Expected key1 to be evicted, but it's still in the cache: %v", value)
	}

	value, ok = cache.Get("key2")
	if !ok {
		t.Errorf("Expected key2 to still be in the cache, but it's not")
	}

	value, ok = cache.Get("key3")
	if !ok {
		t.Errorf("Expected key3 to still be in the cache, but it's not")
	}
}
