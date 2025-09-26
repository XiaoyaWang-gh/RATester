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
		name: "testName",
	}
	name, typeName, spanKind := fa.GetTracingInformation()
	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}
	if typeName != "typeNameForward" {
		t.Errorf("Expected typeName to be 'typeNameForward', but got %s", typeName)
	}
	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be 'trace.SpanKindInternal', but got %s", spanKind)
	}
}
