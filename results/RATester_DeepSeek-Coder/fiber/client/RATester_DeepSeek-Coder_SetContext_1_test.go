package client

import (
	"context"
	"fmt"
	"testing"
)

func TestSetContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &Request{}
	r.SetContext(ctx)

	if r.ctx != ctx {
		t.Errorf("Expected context to be %v, got %v", ctx, r.ctx)
	}
}
