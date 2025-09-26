package xlog

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestFromContextSafe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	xl := FromContextSafe(ctx)
	assert.NotNil(t, xl)

	ctx = context.WithValue(ctx, xlogKey, nil)
	xl = FromContextSafe(ctx)
	assert.NotNil(t, xl)

	ctx = context.WithValue(ctx, xlogKey, &Logger{})
	xl = FromContextSafe(ctx)
	assert.NotNil(t, xl)
}
