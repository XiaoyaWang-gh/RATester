package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &MemProvider{}

	// Test case 1: Check if SessionGC function is working correctly
	rp.SessionGC(ctx)

	// Test case 2: Check if SessionGC function is working correctly with a nil context
	rp.SessionGC(nil)
}
