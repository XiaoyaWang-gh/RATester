package opentelemetry

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func TestNewFilterChainBuilder_2(t *testing.T) {
	type args struct {
		options []FilterChainOption
	}
	tests := []struct {
		name string
		args args
		want *FilterChainBuilder
	}{
		{
			name: "test case 1",
			args: args{
				options: []FilterChainOption{
					func(fcv *FilterChainBuilder) {
						fcv.customSpanFunc = func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
							// do something
						}
					},
				},
			},
			want: &FilterChainBuilder{
				customSpanFunc: func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
					// do something
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewFilterChainBuilder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterChainBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
