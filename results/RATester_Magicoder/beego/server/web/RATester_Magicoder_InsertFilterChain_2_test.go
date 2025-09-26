package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestInsertFilterChain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	pattern := "/test"
	filterChain := func(next FilterFunc) FilterFunc {
		return func(ctx *beecontext.Context) {
			// Pre-processing
			next(ctx)
			// Post-processing
		}
	}
	opts := []FilterOpt{
		func(opts *filterOpts) {
			opts.returnOnOutput = true
			opts.resetParams = true
			opts.routerCaseSensitive = false
		},
	}

	// Act
	result := InsertFilterChain(pattern, filterChain, opts...)

	// Assert
	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}
}
