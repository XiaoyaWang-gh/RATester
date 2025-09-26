package passtlsclientcert

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &passTLSClientCert{}
	name, kind, spanKind := p.GetTracingInformation()
	assert.Equal(t, name, "")
	assert.Equal(t, kind, "")
	assert.Equal(t, spanKind, trace.SpanKindInternal)
}
