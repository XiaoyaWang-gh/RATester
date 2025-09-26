package circuitbreaker

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &circuitBreaker{
		name: "test",
	}
	name, typeName, spanKind := c.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "circuitBreaker", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
