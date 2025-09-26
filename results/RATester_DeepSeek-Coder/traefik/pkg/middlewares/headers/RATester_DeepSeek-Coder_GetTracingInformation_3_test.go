package headers

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &headers{
		name:    "test",
		handler: nil,
	}

	name, typeName, kind := h.GetTracingInformation()

	if name != "test" {
		t.Errorf("Expected name to be 'test', got %s", name)
	}

	if typeName != "headers" {
		t.Errorf("Expected typeName to be 'headers', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, got %v", kind)
	}
}
