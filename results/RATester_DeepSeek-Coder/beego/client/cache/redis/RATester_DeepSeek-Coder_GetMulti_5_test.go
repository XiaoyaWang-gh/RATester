package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestGetMulti_5(t *testing.T) {
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
	keys := []string{"key1", "key2", "key3"}
	values := []interface{}{"value1", "value2", "value3"}

	// Setup test data
	conn := rc.p.Get()
	for i, key := range keys {
		_, err := conn.Do("SET", key, values[i])
		if err != nil {
			t.Fatalf("Failed to set test data: %v", err)
		}
	}
	err := conn.Close()
	if err != nil {
		t.Fatalf("Failed to close connection: %v", err)
	}

	// Test GetMulti
	result, err := rc.GetMulti(ctx, keys)
	if err != nil {
		t.Errorf("GetMulti failed: %v", err)
	}
	if len(result) != len(values) {
		t.Errorf("Expected %d results, got %d", len(values), len(result))
	}
	for i, value := range values {
		if result[i] != value {
			t.Errorf("Expected value %v at index %d, got %v", value, i, result[i])
		}
	}

	// Cleanup test data
	conn = rc.p.Get()
	for _, key := range keys {
		_, err := conn.Do("DEL", key)
		if err != nil {
			t.Fatalf("Failed to delete test data: %v", err)
		}
	}
	err = conn.Close()
	if err != nil {
		t.Fatalf("Failed to close connection: %v", err)
	}
}
