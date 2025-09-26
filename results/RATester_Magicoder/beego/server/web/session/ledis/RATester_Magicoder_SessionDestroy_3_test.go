package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	lp := &Provider{
		maxlifetime: 100,
		SavePath:    "/tmp/sessions",
		Db:          0,
	}

	sid := "test_session"

	err := lp.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy failed: %v", err)
	}
}
