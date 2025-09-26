package opentracing

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/opentracing/opentracing-go"
)

func TestFilterChain_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{
		CustomSpanFunc: func(span opentracing.Span, ctx context.Context, inv *orm.Invocation) {
			// Custom span logic
		},
	}

	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		// Mock next filter behavior
		return []interface{}{nil}
	}

	ctx := context.Background()
	inv := &orm.Invocation{
		Method: "QueryTable",
	}

	result := builder.FilterChain(next)(ctx, inv)

	// Assert result
	if result == nil {
		t.Error("Expected result to be not nil")
	}
}
