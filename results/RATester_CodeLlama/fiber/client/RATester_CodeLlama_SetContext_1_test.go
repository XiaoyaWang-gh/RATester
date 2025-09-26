package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	r := &Request{}
	ctx := context.Background()
	// WHEN
	r.SetContext(ctx)
	// THEN
	assert.Equal(t, ctx, r.Context())
}
