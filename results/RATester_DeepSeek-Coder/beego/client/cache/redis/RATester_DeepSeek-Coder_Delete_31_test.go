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
	_, err := rc.do("SET", key, "test_value")
	if err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}

	err = rc.Delete(ctx, key)
	if err != nil {
		t.Errorf("Failed to delete key: %v", err)
	}

	val, err := redis.String(rc.do("GET", key))
	if err != nil && err != redis.ErrNil {
		t.Errorf("Failed to get key: %v", err)
	}
	if val != "" {
		t.Errorf("Expected empty value, got: %v", val)
	}
}
