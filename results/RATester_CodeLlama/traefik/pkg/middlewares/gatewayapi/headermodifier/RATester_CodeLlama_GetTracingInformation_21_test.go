package headermodifier

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &responseHeaderModifier{
		name: "responseHeaderModifier",
	}
	name, typ, kind := r.GetTracingInformation()
	assert.Equal(t, "responseHeaderModifier", name)
	assert.Equal(t, "responseHeaderModifier", typ)
	assert.Equal(t, trace.SpanKindUnspecified, kind)
}
