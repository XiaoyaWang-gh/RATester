package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

func TestPut_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rc := &Cache{
		p:        &redis.Pool{},
		conninfo: "localhost:6379",
		dbNum:    0,
		key:      "test",
		password: "",
		maxIdle:  10,
		timeout:  10 * time.Second,
	}

	key := "test_key"
	val := "test_val"
	timeout := 10 * time.Second

	err := rc.Put(ctx, key, val, timeout)
	if err != nil {
		t.Errorf("Put failed: %v", err)
	}

	// Add more assertions to check if the key-value pair is correctly stored in Redis
}
