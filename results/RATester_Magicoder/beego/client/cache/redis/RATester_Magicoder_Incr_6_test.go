package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestIncr_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rc := &Cache{
		p: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", ":6379")
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}

	// Setup
	key := "test_key"
	rc.do("SET", key, 1)

	// Test
	err := rc.Incr(ctx, key)
	if err != nil {
		t.Errorf("Incr failed: %v", err)
	}

	// Verify
	val, err := redis.Int(rc.do("GET", key))
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Cleanup
	rc.do("DEL", key)
}
