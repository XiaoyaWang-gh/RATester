package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp",
		Pool:        "default",
		Bucket:      "default",
	}

	provider.SessionGC(ctx)

	// Add assertions here to verify the behavior of the method
}
