package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestDelete_31(t *testing.T) {
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

	key := "test_key"
	value := "test_value"

	// Setup
	err := rc.Put(ctx, key, value, 0)
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	// Test
	err = rc.Delete(ctx, key)
	if err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}

	// Verify
	_, err = rc.Get(ctx, key)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
}
