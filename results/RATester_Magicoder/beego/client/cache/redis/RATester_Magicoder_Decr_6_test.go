package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestDecr_6(t *testing.T) {
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
	key := "test_key"
	rc.do("SET", key, 10)

	// Test Decr
	err := rc.Decr(ctx, key)
	if err != nil {
		t.Errorf("Decr failed: %v", err)
	}

	// Verify result
	val, err := redis.Int(rc.do("GET", key))
	if err != nil {
		t.Errorf("Failed to get value: %v", err)
	}
	if val != 9 {
		t.Errorf("Expected value 9, got %d", val)
	}

	// Cleanup
	rc.do("DEL", key)
}
