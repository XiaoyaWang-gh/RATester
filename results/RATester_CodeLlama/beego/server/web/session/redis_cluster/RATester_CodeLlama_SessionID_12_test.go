package redis_cluster

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &SessionStore{}
	rs.sid = "123456"
	ctx := context.Background()
	if rs.SessionID(ctx) != "123456" {
		t.Error("SessionID failed")
	}
}
