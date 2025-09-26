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
		return []interface{}{nil}
	}

	filter := builder.FilterChain(next)
	ctx := context.Background()
	inv := &orm.Invocation{}
	res := filter(ctx, inv)

	if len(res) != 1 || res[0] != nil {
		t.Errorf("Expected result to be [nil], but got %v", res)
	}
}
