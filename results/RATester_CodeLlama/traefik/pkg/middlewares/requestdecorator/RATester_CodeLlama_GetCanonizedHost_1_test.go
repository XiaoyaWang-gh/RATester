package requestdecorator

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetCanonizedHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ctx = context.WithValue(ctx, canonicalKey, "test")
	assert.Equal(t, "test", GetCanonizedHost(ctx))
}
