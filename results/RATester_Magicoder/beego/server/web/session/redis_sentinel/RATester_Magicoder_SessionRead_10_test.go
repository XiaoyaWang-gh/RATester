package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/server/web/session"
	"github.com/redis/go-redis/v9"
)

func TestSessionRead_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		maxlifetime: 100,
		poollist:    redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}

	// Test case 1: Session does not exist
	sid := "non-existent-session"
	_, err := rp.SessionRead(ctx, sid)
	if err == nil {
		t.Errorf("Expected error for non-existent session, but got nil")
	}

	// Test case 2: Session exists but is empty
	sid = "empty-session"
	rp.poollist.Set(ctx, sid, "", time.Duration(rp.maxlifetime)*time.Second)
	_, err = rp.SessionRead(ctx, sid)
	if err != nil {
		t.Errorf("Unexpected error for empty session: %v", err)
	}

	// Test case 3: Session exists and is not empty
	sid = "non-empty-session"
	kv := map[interface{}]interface{}{"key": "value"}
	data, _ := session.EncodeGob(kv)
	rp.poollist.Set(ctx, sid, string(data), time.Duration(rp.maxlifetime)*time.Second)
	_, err = rp.SessionRead(ctx, sid)
	if err != nil {
		t.Errorf("Unexpected error for non-empty session: %v", err)
	}
}
