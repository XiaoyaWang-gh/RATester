package redirect

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &redirect{
		name: "test_name",
	}
	name, typeName, kind := r.GetTracingInformation()
	if name != "test_name" {
		t.Errorf("Expected name to be 'test_name', but got %s", name)
	}
	if typeName != "redirect" {
		t.Errorf("Expected typeName to be 'redirect', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
