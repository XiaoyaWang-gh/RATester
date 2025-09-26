package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestIncr_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rc := &Cache{
		conn:     &ssdb.Client{},
		conninfo: []string{},
	}

	err := rc.Incr(ctx, "test_key")
	if err != nil {
		t.Errorf("Incr failed: %v", err)
	}
}
