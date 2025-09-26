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
		name: "testName",
	}
	name, typeName, kind := wl.GetTracingInformation()
	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}
	if typeName != "ipWhiteLister" {
		t.Errorf("Expected typeName to be 'ipWhiteLister', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
