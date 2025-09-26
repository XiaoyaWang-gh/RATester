package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func TestBuildSpan_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		span otelTrace.Span
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
			builder.buildSpan(tt.args.ctx, tt.args.span, tt.args.inv)
		})
	}
}
