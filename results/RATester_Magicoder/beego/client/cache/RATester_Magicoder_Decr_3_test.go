package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestDecr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := &MemoryCache{
		items: make(map[string]*MemoryItem),
	}

	key := "testKey"
	val := "10"
	cache.items[key] = &MemoryItem{
		val: val,
	}

	err := cache.Decr(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expectedVal := "9"
	if cache.items[key].val != expectedVal {
		t.Errorf("Expected value to be %v, but got %v", expectedVal, cache.items[key].val)
	}

	key = "nonExistentKey"
	err = cache.Decr(ctx, key)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
