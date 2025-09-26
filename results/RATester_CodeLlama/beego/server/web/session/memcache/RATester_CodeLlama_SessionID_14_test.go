package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &SessionStore{}
	rs.sid = "test"
	ctx := context.Background()
	if rs.SessionID(ctx) != "test" {
		t.Error("SessionID() failed")
	}
}
