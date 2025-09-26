package stripprefixregex

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &stripPrefixRegex{}
	name, typeName, spanKind := s.GetTracingInformation()
	assert.Equal(t, name, "")
	assert.Equal(t, typeName, "stripPrefixRegex")
	assert.Equal(t, spanKind, trace.SpanKindInternal)
}
