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
		name:   "testBuffer",
		buffer: &oxybuffer.Buffer{},
	}

	name, typeName, kind := b.GetTracingInformation()

	if name != "testBuffer" {
		t.Errorf("Expected name to be 'testBuffer', got %s", name)
	}

	if typeName != "*oxybuffer.Buffer" {
		t.Errorf("Expected typeName to be '*oxybuffer.Buffer', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, got %v", kind)
	}
}
