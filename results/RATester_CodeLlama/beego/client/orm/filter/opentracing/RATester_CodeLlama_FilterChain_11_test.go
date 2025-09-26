package opentracing

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"gotest.tools/assert"
)

func TestFilterChain_11(t *testing.T) {
	t.Run("test FilterChain", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		builder := &FilterChainBuilder{}
		next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
			return []interface{}{}
		}
		res := builder.FilterChain(next)(context.Background(), &orm.Invocation{})
		assert.Equal(t, []interface{}{}, res)
	})
}
