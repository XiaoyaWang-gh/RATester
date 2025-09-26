package net

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

	ctx := context.Background()
	conn := &ContextConn{ctx: ctx}

	if conn.Context() != ctx {
		t.Errorf("Expected context to be %v, but got %v", ctx, conn.Context())
	}
}
