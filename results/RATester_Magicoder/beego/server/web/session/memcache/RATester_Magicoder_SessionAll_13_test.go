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
	result := rp.SessionAll(ctx)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
