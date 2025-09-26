package redirect

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := redirect{
		name:       "testRedirect",
		next:       nil,
		scheme:     nil,
		hostname:   nil,
		port:       nil,
		path:       nil,
		pathPrefix: nil,
		statusCode: 302,
	}

	name, typeName, kind := r.GetTracingInformation()

	if name != "testRedirect" {
		t.Errorf("Expected name to be 'testRedirect', but got %s", name)
	}

	if typeName != "redirect" {
		t.Errorf("Expected typeName to be 'redirect', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %s", kind)
	}
}
