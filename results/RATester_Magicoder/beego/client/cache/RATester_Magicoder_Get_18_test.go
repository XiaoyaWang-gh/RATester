package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	ctx := context.Background()

	key := "test_key"
	val := "test_value"

	// Setup
	err := fc.Put(ctx, key, val, 0)
	if err != nil {
		t.Fatalf("Failed to put value: %v", err)
	}

	// Test
	got, err := fc.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if got != val {
		t.Fatalf("Expected %v, got %v", val, got)
	}

	// Cleanup
	err = fc.Delete(ctx, key)
	if err != nil {
		t.Fatalf("Failed to delete value: %v", err)
	}
}
