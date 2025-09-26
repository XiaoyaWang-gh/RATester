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

	r := &Request{}
	ctx := r.Context()
	if ctx != context.Background() {
		t.Errorf("Context() = %v, want %v", ctx, context.Background())
	}
}
