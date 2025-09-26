package net

import (
	"context"
	"fmt"
	"testing"
)

func TestWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	conn := &ContextConn{}
	conn.WithContext(ctx)
	if conn.ctx != ctx {
		t.Errorf("conn.ctx = %v, want %v", conn.ctx, ctx)
	}
}
