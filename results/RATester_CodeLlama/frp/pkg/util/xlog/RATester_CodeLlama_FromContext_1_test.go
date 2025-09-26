package xlog

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestFromContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	xl, ok := FromContext(ctx)
	assert.False(t, ok)
	assert.Nil(t, xl)

	ctx = context.WithValue(ctx, xlogKey, &Logger{})
	xl, ok = FromContext(ctx)
	assert.True(t, ok)
	assert.NotNil(t, xl)
}
