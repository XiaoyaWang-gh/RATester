package ipwhitelist

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	wl := &ipWhiteLister{
		name: "test",
	}

	name, typeName, kind := wl.GetTracingInformation()

	if name != "test" {
		t.Errorf("Expected name to be 'test', got %s", name)
	}

	if typeName != "ipWhiteLister" {
		t.Errorf("Expected typeName to be 'ipWhiteLister', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, got %v", kind)
	}
}
