package prometheus

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestFilterChain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{
		AppName:    "testApp",
		ServerName: "testServer",
		RunMode:    "testMode",
	}

	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		// Mock the next filter
		return []interface{}{nil}
	}

	filter := builder.FilterChain(next)

	ctx := context.Background()
	inv := &orm.Invocation{
		// Mock the invocation
	}

	res := filter(ctx, inv)

	// Assert the result
	if res == nil {
		t.Error("Expected result to be not nil")
	}
}
