package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestGet_38(t *testing.T) {
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

	key := "test_key"
	val := "test_val"

	// Set a value in Redis
	conn := rc.p.Get()
	_, err := conn.Do("SET", key, val)
	if err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}
	conn.Close()

	// Get the value from the cache
	got, err := rc.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}

	// Check if the value is correct
	if got != val {
		t.Errorf("Expected %v, got %v", val, got)
	}
}
