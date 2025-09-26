package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestClearAll_4(t *testing.T) {
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
	err := rc.ClearAll(ctx)
	if err != nil {
		t.Errorf("ClearAll failed: %v", err)
	}
}
