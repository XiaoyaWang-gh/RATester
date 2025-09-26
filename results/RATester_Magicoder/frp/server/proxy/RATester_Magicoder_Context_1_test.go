package proxy

import (
	"context"
	"fmt"
	"testing"
)

func TestContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		ctx: context.Background(),
	}

	ctx := pxy.Context()

	if ctx != pxy.ctx {
		t.Errorf("Expected %v, got %v", pxy.ctx, ctx)
	}
}
