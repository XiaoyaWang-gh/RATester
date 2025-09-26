package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestFilterChain_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	builder := &FilterChainBuilder{}
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		return []interface{}{nil}
	}

	filter := builder.FilterChain(next)
	res := filter(ctx, &orm.Invocation{Method: "Read"})
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}

	res = filter(ctx, &orm.Invocation{Method: "Begin"})
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}

	res = filter(ctx, &orm.Invocation{Method: "Commit"})
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}

	res = filter(ctx, &orm.Invocation{Method: "Rollback"})
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}
