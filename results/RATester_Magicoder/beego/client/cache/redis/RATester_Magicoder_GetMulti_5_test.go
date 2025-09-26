package redis

import (
	"context"
	"fmt"
	"reflect"
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
				c, err := redis.Dial("tcp", ":6379")
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}

	keys := []string{"key1", "key2", "key3"}
	values := []interface{}{"value1", "value2", "value3"}

	// Setup test data
	for i, key := range keys {
		if err := rc.Put(ctx, key, values[i], 0); err != nil {
			t.Fatalf("Failed to setup test data: %v", err)
		}
	}

	// Test GetMulti
	got, err := rc.GetMulti(ctx, keys)
	if err != nil {
		t.Fatalf("GetMulti failed: %v", err)
	}

	// Verify results
	for i, key := range keys {
		if !reflect.DeepEqual(got[i], values[i]) {
			t.Errorf("GetMulti(%q) = %v, want %v", key, got[i], values[i])
		}
	}

	// Cleanup
	for _, key := range keys {
		if err := rc.Delete(ctx, key); err != nil {
			t.Fatalf("Failed to cleanup test data: %v", err)
		}
	}
}
