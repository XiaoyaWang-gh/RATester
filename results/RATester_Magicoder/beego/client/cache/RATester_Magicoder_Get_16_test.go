package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGet_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := &MemoryCache{
		dur:   1 * time.Minute,
		items: make(map[string]*MemoryItem),
	}

	// Test case 1: key exists and not expired
	cache.items["key1"] = &MemoryItem{
		val:         "value1",
		createdTime: time.Now(),
		lifespan:    2 * time.Minute,
	}
	val, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected value 'value1', got '%v'", val)
	}

	// Test case 2: key exists but expired
	cache.items["key2"] = &MemoryItem{
		val:         "value2",
		createdTime: time.Now().Add(-3 * time.Minute),
		lifespan:    2 * time.Minute,
	}
	val, err = cache.Get(ctx, "key2")
	if err != ErrKeyExpired {
		t.Errorf("Expected error ErrKeyExpired, got '%v'", err)
	}
	if val != nil {
		t.Errorf("Expected nil value, got '%v'", val)
	}

	// Test case 3: key does not exist
	val, err = cache.Get(ctx, "key3")
	if err != ErrKeyNotExist {
		t.Errorf("Expected error ErrKeyNotExist, got '%v'", err)
	}
	if val != nil {
		t.Errorf("Expected nil value, got '%v'", val)
	}
}
