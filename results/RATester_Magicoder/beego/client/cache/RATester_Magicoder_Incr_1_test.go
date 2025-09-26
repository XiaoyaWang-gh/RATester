package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIncr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cache := NewMemoryCache()

	// Test with a key that does not exist
	err := cache.Incr(ctx, "nonExistentKey")
	if err != ErrKeyNotExist {
		t.Errorf("Expected ErrKeyNotExist, got %v", err)
	}

	// Test with a key that exists
	cache.Put(ctx, "existingKey", 1, time.Hour)
	err = cache.Incr(ctx, "existingKey")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	val, _ := cache.Get(ctx, "existingKey")
	if val.(int) != 2 {
		t.Errorf("Expected value to be 2, got %v", val)
	}

	// Test with a key that exists and has a non-integer value
	cache.Put(ctx, "nonIntKey", "notAnInt", time.Hour)
	err = cache.Incr(ctx, "nonIntKey")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
