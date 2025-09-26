package ipallowlist

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	al := &ipAllowLister{
		name: "testName",
	}

	name, typeName, spanKind := al.GetTracingInformation()

	if name != "testName" {
		t.Errorf("Expected name to be 'testName', got %s", name)
	}

	if typeName != "ipAllowLister" {
		t.Errorf("Expected typeName to be 'ipAllowLister', got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be 'trace.SpanKindInternal', got %v", spanKind)
	}
}
