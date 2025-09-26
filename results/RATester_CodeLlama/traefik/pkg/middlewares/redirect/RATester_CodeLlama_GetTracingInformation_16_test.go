package redirect

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &redirect{
		name: "test",
	}
	name, typeName, spanKind := r.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "redirect", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
