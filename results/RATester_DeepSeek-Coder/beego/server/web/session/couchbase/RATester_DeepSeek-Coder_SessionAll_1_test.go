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
	provider := &Provider{
		maxlifetime: 100,
		SavePath:    "/tmp",
		Pool:        "default",
		Bucket:      "default",
	}

	sessionAll := provider.SessionAll(ctx)
	if sessionAll != 0 {
		t.Errorf("Expected 0, got %d", sessionAll)
	}
}
