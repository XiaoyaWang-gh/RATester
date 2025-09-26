package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestIsExist_5(t *testing.T) {
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

	// Setup test data
	key := "test_key"
	value := "test_value"
	rc.Put(ctx, key, value, 0)

	// Test existing key
	exist, err := rc.IsExist(ctx, key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exist {
		t.Errorf("Expected key to exist, but it does not")
	}

	// Test non-existing key
	exist, err = rc.IsExist(ctx, "non_existing_key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if exist {
		t.Errorf("Expected key to not exist, but it does")
	}

	// Cleanup test data
	rc.Delete(ctx, key)
}
