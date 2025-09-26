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
		p: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", ":6379")
			},
		},
		conninfo: ":6379",
		dbNum:    0,
		key:      "test",
		password: "",
		maxIdle:  10,
		timeout:  1 * time.Second,
	}

	key := "test_key"
	val := "test_val"
	timeout := 1 * time.Second

	err := rc.Put(ctx, key, val, timeout)
	if err != nil {
		t.Errorf("Put failed: %v", err)
	}

	conn := rc.p.Get()
	defer conn.Close()

	got, err := redis.String(conn.Do("GET", key))
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}

	if got != val {
		t.Errorf("Expected %v, got %v", val, got)
	}
}
