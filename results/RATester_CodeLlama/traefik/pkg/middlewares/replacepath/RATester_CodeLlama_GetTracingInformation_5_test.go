package replacepath

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &replacePath{
		name: "test",
	}
	name, typeName, spanKind := r.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "replacePath", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
