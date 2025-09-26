package headers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &headers{
		name: "test",
	}
	name, typeName, spanKind := h.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "headers", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
