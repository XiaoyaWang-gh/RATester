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

	ctx := context.Background()
	req := &Request{}

	req.SetContext(ctx)

	if req.ctx != ctx {
		t.Errorf("Expected ctx to be %v, but got %v", ctx, req.ctx)
	}
}
