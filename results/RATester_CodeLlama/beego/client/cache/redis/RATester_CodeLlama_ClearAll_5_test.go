package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestClearAll_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{
		p: &redis.Pool{},
	}
	err := rc.ClearAll(context.Background())
	if err != nil {
		t.Errorf("ClearAll() error = %v", err)
		return
	}
}
