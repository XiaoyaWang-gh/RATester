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

	builder := &FilterChainBuilder{}
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		return []interface{}{nil}
	}
	filter := builder.FilterChain(next)
	ctx := context.Background()
	inv := &orm.Invocation{Method: "Begin"}
	res := filter(ctx, inv)
	if res[0] != nil {
		t.Errorf("Expected nil, got %v", res[0])
	}
	inv = &orm.Invocation{Method: "Commit"}
	res = filter(ctx, inv)
	if res[0] != nil {
		t.Errorf("Expected nil, got %v", res[0])
	}
	inv = &orm.Invocation{Method: "Rollback"}
	res = filter(ctx, inv)
	if res[0] != nil {
		t.Errorf("Expected nil, got %v", res[0])
	}
	inv = &orm.Invocation{Method: "Other"}
	res = filter(ctx, inv)
	if res[0] != nil {
		t.Errorf("Expected nil, got %v", res[0])
	}
}
