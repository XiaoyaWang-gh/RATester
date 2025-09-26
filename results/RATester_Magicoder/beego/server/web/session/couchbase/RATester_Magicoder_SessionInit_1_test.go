package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	maxlifetime := int64(3600)
	cfg := `{"save_path":"/tmp","pool":"default","bucket":"test"}`

	cp := &Provider{}
	err := cp.SessionInit(ctx, maxlifetime, cfg)
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if cp.maxlifetime != maxlifetime {
		t.Errorf("SessionInit failed: maxlifetime not set correctly")
	}

	if cp.SavePath != "/tmp" || cp.Pool != "default" || cp.Bucket != "test" {
		t.Errorf("SessionInit failed: cfg not set correctly")
	}
}
