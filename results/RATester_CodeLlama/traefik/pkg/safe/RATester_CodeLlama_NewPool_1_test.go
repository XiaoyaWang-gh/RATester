package safe

import (
	"context"
	"fmt"
	"testing"
)

func TestNewPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pool := NewPool(ctx)
	if pool.ctx != ctx {
		t.Errorf("pool.ctx = %v, want %v", pool.ctx, ctx)
	}
	if pool.cancel == nil {
		t.Errorf("pool.cancel = nil, want not nil")
	}
}
