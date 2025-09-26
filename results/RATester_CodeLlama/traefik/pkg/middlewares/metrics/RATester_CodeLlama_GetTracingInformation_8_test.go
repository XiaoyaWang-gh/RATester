package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &metricsMiddleware{}
	name, typeName, spanKind := m.GetTracingInformation()
	assert.Equal(t, "", name)
	assert.Equal(t, "", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
