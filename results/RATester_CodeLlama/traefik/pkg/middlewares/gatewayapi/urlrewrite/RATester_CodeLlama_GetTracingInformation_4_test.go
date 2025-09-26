package urlrewrite

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	u := urlRewrite{
		name: "test",
	}
	name, typeName, spanKind := u.GetTracingInformation()
	assert.Equal(t, "test", name)
	assert.Equal(t, "urlRewrite", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
