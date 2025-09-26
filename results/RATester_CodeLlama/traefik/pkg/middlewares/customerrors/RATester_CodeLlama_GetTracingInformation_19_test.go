package customerrors

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &customErrors{
		name: "test",
	}
	name, typeName, spanKind := c.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "customErrors", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
