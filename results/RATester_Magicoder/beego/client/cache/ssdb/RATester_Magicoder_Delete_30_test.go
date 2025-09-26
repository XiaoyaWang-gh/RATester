package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestDelete_30(t *testing.T) {
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

	err := rc.Delete(ctx, "test_key")
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}
