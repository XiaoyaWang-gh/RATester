package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDecr_2(t *testing.T) {
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
	val := "10"

	// Setup
	err := fc.Put(context.Background(), key, val, time.Duration(fc.EmbedExpiry))
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	// Test
	err = fc.Decr(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to Decr: %v", err)
	}

	// Verify
	data, err := fc.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to Get: %v", err)
	}

	if data.(string) != "9" {
		t.Fatalf("Expected 9, got %v", data)
	}

	// Cleanup
	err = fc.Delete(context.Background(), key)
	if err != nil {
		t.Fatalf("Failed to cleanup: %v", err)
	}
}
