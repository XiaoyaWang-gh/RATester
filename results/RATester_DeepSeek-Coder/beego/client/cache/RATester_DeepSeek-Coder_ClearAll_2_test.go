package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestClearAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := &MemoryCache{
		items: map[string]*MemoryItem{
			"key1": {val: "value1"},
			"key2": {val: "value2"},
		},
	}

	err := cache.ClearAll(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(cache.items) != 0 {
		t.Errorf("Expected items to be empty, got %v", cache.items)
	}
}
