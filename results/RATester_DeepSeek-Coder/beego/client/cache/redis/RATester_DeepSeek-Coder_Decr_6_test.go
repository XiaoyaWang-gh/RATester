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

	key := "test_key"
	value := "100"

	// Set initial value
	rc.do("SET", key, value)

	// Decrement value
	err := rc.Decr(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if value was decremented
	result, _ := redis.String(rc.do("GET", key))
	if result != "99" {
		t.Errorf("Expected value to be 99, got %s", result)
	}
}
