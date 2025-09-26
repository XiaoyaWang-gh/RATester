package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/stretchr/testify/assert"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func TestWithCustomSpanFunc_1(t *testing.T) {
	t.Parallel()

	t.Run("should set customSpanFunc", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		fcv := &FilterChainBuilder{}
		customSpanFunc := func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {}

		WithCustomSpanFunc(customSpanFunc)(fcv)

		assert.Equal(t, customSpanFunc, fcv.customSpanFunc)
	})
}
