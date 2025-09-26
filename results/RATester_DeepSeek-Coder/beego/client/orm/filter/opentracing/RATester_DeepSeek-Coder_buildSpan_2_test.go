package opentracing

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/opentracing/opentracing-go"
)

func TestBuildSpan_2(t *testing.T) {
	type args struct {
		span opentracing.Span
		ctx  context.Context
		inv  *orm.Invocation
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			builder := &FilterChainBuilder{}
			builder.buildSpan(tt.args.span, tt.args.ctx, tt.args.inv)
		})
	}
}
