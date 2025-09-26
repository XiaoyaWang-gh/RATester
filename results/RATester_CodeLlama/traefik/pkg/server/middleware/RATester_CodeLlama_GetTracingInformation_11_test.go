package middleware

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &traceablePlugin{
		name: "test",
	}
	name, typeName, spanKind := s.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "traceablePlugin", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
