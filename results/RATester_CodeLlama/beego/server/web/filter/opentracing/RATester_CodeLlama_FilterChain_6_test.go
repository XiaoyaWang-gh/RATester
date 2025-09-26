package opentracing

import (
	"fmt"
	"testing"

	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func TestFilterChain_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &FilterChainBuilder{}
	next := func(ctx *beegoCtx.Context) {}
	filter := builder.FilterChain(next)
	filter(nil)
}
