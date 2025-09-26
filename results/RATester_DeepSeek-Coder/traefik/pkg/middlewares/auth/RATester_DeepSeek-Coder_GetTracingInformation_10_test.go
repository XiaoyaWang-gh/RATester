package auth

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fa := &forwardAuth{
		name: "testForwardAuth",
	}

	name, typeName, spanKind := fa.GetTracingInformation()

	if name != "testForwardAuth" {
		t.Errorf("Expected name to be 'testForwardAuth', got %s", name)
	}

	if typeName != "forwardAuth" {
		t.Errorf("Expected typeName to be 'forwardAuth', got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be %v, got %v", trace.SpanKindInternal, spanKind)
	}
}
