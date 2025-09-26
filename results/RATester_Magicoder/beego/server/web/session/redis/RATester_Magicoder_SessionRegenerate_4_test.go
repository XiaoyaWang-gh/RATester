package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionRegenerate_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		maxlifetime: 10,
		poollist:    redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}

	// Test case 1: oldsid does not exist
	oldsid := "oldsid"
	sid := "sid"
	_, err := rp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: oldsid exists
	rp.poollist.Set(ctx, oldsid, "value", 0)
	_, err = rp.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Clean up
	rp.poollist.Del(ctx, oldsid)
	rp.poollist.Del(ctx, sid)
}
