package addprefix

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &addPrefix{
		name: "test",
	}
	name, typeName, spanKind := a.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "addPrefix", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
