package client

import (
	"context"
	"fmt"
	"testing"
)

func TestContext_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{}
	ctx := req.Context()
	if ctx != context.Background() {
		t.Errorf("Expected context.Background(), got %v", ctx)
	}

	ctx = context.WithValue(context.Background(), "key", "value")
	req.ctx = ctx
	if req.Context() != ctx {
		t.Errorf("Expected %v, got %v", ctx, req.Context())
	}
}
