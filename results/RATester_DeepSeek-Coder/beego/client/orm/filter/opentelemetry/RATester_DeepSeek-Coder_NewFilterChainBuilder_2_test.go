package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func TestNewFilterChainBuilder_2(t *testing.T) {
	t.Run("should return a FilterChainBuilder with custom options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		customOption := func(fcv *FilterChainBuilder) {
			fcv.customSpanFunc = func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
				// custom logic here
			}
		}

		builder := NewFilterChainBuilder(customOption)

		if builder.customSpanFunc == nil {
			t.Errorf("expected customSpanFunc to be set, but it was nil")
		}
	})
}
