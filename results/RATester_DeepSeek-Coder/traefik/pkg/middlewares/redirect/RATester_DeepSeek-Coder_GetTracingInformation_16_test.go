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
		name: "testRedirect",
	}

	name, typeName, kind := r.GetTracingInformation()

	if name != "testRedirect" {
		t.Errorf("Expected name to be 'testRedirect', got %s", name)
	}

	if typeName != "redirect" {
		t.Errorf("Expected typeName to be 'redirect', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be %v, got %v", trace.SpanKindInternal, kind)
	}
}
