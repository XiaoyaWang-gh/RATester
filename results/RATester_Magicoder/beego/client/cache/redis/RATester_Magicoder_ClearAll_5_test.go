package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestClearAll_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rc := &Cache{
		p: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", ":6379")
			},
		},
	}

	// Setup test data
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		if err := rc.Put(ctx, key, "value", 0); err != nil {
			t.Fatalf("Failed to setup test data: %v", err)
		}
	}

	// Test ClearAll
	if err := rc.ClearAll(ctx); err != nil {
		t.Fatalf("Failed to clear all keys: %v", err)
	}

	// Verify that all keys have been deleted
	for _, key := range keys {
		_, err := rc.Get(ctx, key)
		if err == nil {
			t.Errorf("Key %s was not deleted", key)
		}
	}
}
