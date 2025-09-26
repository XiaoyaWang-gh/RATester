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
	name, typeName, kind := al.GetTracingInformation()

	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}

	if typeName != "" {
		t.Errorf("Expected typeName to be '', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
