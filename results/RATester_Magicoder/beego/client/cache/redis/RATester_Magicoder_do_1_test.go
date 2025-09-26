package redis

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func Testdo_1(t *testing.T) {
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

	_, err := rc.do("SET", "key", "value")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = rc.do("GET", "key")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = rc.do("DEL", "key")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
