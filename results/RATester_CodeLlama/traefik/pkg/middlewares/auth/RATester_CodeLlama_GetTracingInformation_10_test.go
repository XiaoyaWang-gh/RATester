package auth

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fa := &forwardAuth{
		name: "test",
	}
	name, kind, spanKind := fa.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "forwardAuth", kind)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
