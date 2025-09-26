package buffering

import (
	"fmt"
	"testing"

	oxybuffer "github.com/vulcand/oxy/v2/buffer"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &buffer{
		name:   "test_buffer",
		buffer: &oxybuffer.Buffer{},
	}

	name, typeName, kind := b.GetTracingInformation()

	if name != "test_buffer" {
		t.Errorf("Expected name to be 'test_buffer', but got %s", name)
	}

	if typeName != "buffer" {
		t.Errorf("Expected typeName to be 'buffer', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %s", kind)
	}
}
