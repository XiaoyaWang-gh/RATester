package buffering

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &buffer{
		name:   "test",
		buffer: nil,
	}
	name, typeName, spanKind := b.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "buffer", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
