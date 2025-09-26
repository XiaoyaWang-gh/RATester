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
				c, err := redis.Dial("tcp", ":6379")
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}

	keys, err := rc.Scan("*")
	if err != nil {
		t.Errorf("Scan failed, error: %v", err)
	}

	if len(keys) == 0 {
		t.Errorf("Scan failed, no keys found")
	}

	for _, key := range keys {
		if key == "" {
			t.Errorf("Scan failed, empty key found")
		}
	}
}
