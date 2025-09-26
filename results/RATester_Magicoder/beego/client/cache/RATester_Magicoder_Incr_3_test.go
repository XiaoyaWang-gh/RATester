package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIncr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "/tmp/test",
		FileSuffix:     ".cache",
		DirectoryLevel: 2,
		EmbedExpiry:    3600,
	}

	key := "test_key"
	val := "1"

	// Setup
	err := fc.Put(context.Background(), key, val, time.Duration(fc.EmbedExpiry))
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	// Test
	err = fc.Incr(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to increment value: %v", err)
	}

	// Verify
	result, err := fc.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if result != "2" {
		t.Fatalf("Expected value to be 2, got %v", result)
	}

	// Cleanup
	err = fc.Delete(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to cleanup test: %v", err)
	}
}
