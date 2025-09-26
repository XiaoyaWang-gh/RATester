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
				return redis.Dial("tcp", ":6379")
			},
		},
	}

	key := "test_key"

	// Set initial value to 0
	rc.do("SET", key, 0)

	// Test Incr
	err := rc.Incr(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	val, err := redis.Int(rc.do("GET", key))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected value to be 1, got %v", val)
	}

	// Test error case
	_, err = rc.do("DEL", key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = rc.Incr(ctx, key)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
