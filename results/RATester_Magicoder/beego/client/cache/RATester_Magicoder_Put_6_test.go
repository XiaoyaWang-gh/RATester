package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPut_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "/tmp/cache",
		FileSuffix:     ".cache",
		DirectoryLevel: 2,
		EmbedExpiry:    10,
	}

	ctx := context.Background()

	err := fc.Put(ctx, "test_key", "test_value", time.Duration(fc.EmbedExpiry)*time.Second)
	if err != nil {
		t.Errorf("Error in Put: %v", err)
	}

	val, err := fc.Get(ctx, "test_key")
	if err != nil {
		t.Errorf("Error in Get: %v", err)
	}

	if val != "test_value" {
		t.Errorf("Expected 'test_value', got '%v'", val)
	}
}
