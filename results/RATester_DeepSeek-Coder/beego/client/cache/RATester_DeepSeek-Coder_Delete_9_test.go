package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDelete_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := NewMemoryCache()

	key := "testKey"
	value := "testValue"

	// Add a value to the cache
	err := cache.Put(ctx, key, value, time.Hour)
	if err != nil {
		t.Errorf("Error adding value to cache: %v", err)
	}

	// Delete the value from the cache
	err = cache.Delete(ctx, key)
	if err != nil {
		t.Errorf("Error deleting value from cache: %v", err)
	}

	// Check if the value is deleted
	val, err := cache.Get(ctx, key)
	if err != nil {
		t.Errorf("Error getting value from cache: %v", err)
	}
	if val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}
