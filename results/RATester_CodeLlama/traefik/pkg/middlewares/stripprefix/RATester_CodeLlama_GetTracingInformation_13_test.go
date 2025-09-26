package stripprefix

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &stripPrefix{
		name: "test",
	}
	name, typeName, spanKind := s.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "stripPrefix", typeName)
	assert.Equal(t, trace.SpanKindUnspecified, spanKind)
}
