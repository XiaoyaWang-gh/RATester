package xlog

import (
	"context"
	"fmt"
	"testing"
)

func TestNewContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	xl := &Logger{}
	ctx = NewContext(ctx, xl)
	if ctx == nil {
		t.Error("NewContext return nil")
	}
	if ctx.Value(xlogKey) != xl {
		t.Error("NewContext return wrong value")
	}
}
