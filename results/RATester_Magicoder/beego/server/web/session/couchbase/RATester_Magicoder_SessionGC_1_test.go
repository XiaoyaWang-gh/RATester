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
	cp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp/sessions",
	}

	err := cp.SessionInit(ctx, cp.maxlifetime, cp.SavePath)
	if err != nil {
		t.Fatalf("SessionInit failed: %v", err)
	}

	cp.SessionGC(ctx)

	// Add assertions here to verify the behavior of SessionGC
}
