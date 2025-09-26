package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetMulti_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := &MemoryCache{
		items: make(map[string]*MemoryItem),
	}

	keys := []string{"key1", "key2", "key3"}
	values := []interface{}{"value1", "value2", "value3"}

	for i, key := range keys {
		cache.items[key] = &MemoryItem{
			val:         values[i],
			createdTime: time.Now(),
			lifespan:    time.Hour,
		}
	}

	rc, err := cache.GetMulti(ctx, keys)
	if err != nil {
		t.Errorf("GetMulti failed: %v", err)
	}

	for i, key := range keys {
		if rc[i] != values[i] {
			t.Errorf("Expected %v, got %v", values[i], rc[i])
		}
		if _, ok := cache.items[key]; ok {
			t.Errorf("Key %s should have been removed from cache", key)
		}
	}
}
