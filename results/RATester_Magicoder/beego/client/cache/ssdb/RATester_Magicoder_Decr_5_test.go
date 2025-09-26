package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestDecr_5(t *testing.T) {
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

	err := rc.Decr(ctx, "test_key")
	if err != nil {
		t.Errorf("Decr failed: %v", err)
	}
}
