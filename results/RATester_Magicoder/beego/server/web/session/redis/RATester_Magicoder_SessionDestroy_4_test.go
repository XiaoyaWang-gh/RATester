package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionDestroy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp",
		Poolsize:    10,
		Password:    "password",
		DbNum:       0,
		poollist:    redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}

	sid := "test_session"

	err := rp.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy failed: %v", err)
	}

	_, err = rp.poollist.Get(ctx, sid).Result()
	if err != redis.Nil {
		t.Errorf("SessionDestroy failed: session %s should not exist", sid)
	}
}
