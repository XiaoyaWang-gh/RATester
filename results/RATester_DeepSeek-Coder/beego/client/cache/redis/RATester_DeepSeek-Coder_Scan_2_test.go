package redis

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestScan_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{
		p: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", ":6379")
			},
		},
	}

	keys, err := rc.Scan("*")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(keys) == 0 {
		t.Errorf("Expected at least one key, got 0")
	}
}
