package redis_cluster

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionDestroy_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		poollist: &redis.ClusterClient{},
	}

	sid := "test_sid"

	err := rp.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy failed: %v", err)
	}
}
