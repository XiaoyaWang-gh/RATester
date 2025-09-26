package compress

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &compress{}
	name, typeName, spanKind := c.GetTracingInformation()
	assert.Equal(t, name, "")
	assert.Equal(t, typeName, "compress")
	assert.Equal(t, spanKind, trace.SpanKindInternal)
}
