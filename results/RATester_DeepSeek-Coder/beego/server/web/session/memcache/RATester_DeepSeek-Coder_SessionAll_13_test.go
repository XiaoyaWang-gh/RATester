package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &MemProvider{}
	all := rp.SessionAll(ctx)
	if all != 0 {
		t.Errorf("Expected 0, got %d", all)
	}
}
