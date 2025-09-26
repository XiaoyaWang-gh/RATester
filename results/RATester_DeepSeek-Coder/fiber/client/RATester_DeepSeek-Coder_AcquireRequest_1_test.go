package client

import (
	"context"
	"fmt"
	"testing"
)

func TestAcquireRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	req := AcquireRequest()

	if req == nil {
		t.Fatal("Expected a non-nil request")
	}

	if req.ctx != nil {
		t.Errorf("Expected nil context, got %v", req.ctx)
	}

	req.SetContext(context.Background())

	if req.ctx == nil {
		t.Error("Expected non-nil context after setting it")
	}

	req.Reset()

	if req.ctx != nil {
		t.Error("Expected context to be reset after calling Reset()")
	}
}
