package prometheus

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestFilterChain_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{}
	next := func(ctx *context.Context) {
		// Implement the logic of the next function
	}

	filterChain := builder.FilterChain(next)

	// Test the filterChain function
	ctx := &context.Context{}
	filterChain(ctx)

	// Add assertions to validate the behavior of the function
}
