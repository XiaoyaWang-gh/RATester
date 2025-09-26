package prometheus

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/stretchr/testify/assert"
)

func TestFilterChain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{
		AppName:    "test",
		ServerName: "test",
		RunMode:    "test",
	}

	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		return []interface{}{}
	}

	filter := builder.FilterChain(next)
	ctx := context.Background()
	inv := &orm.Invocation{}
	res := filter(ctx, inv)
	assert.Equal(t, []interface{}{}, res)
}
