package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp/sessions",
		Pool:        "default",
		Bucket:      "sessions",
	}

	result := cp.SessionAll(ctx)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
