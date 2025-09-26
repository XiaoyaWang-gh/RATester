package stripprefixregex

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sr := &stripPrefixRegex{
		name: "test_name",
	}
	name, typeName, kind := sr.GetTracingInformation()
	if name != "test_name" {
		t.Errorf("Expected name to be 'test_name', but got %s", name)
	}
	if typeName != "stripPrefixRegex" {
		t.Errorf("Expected typeName to be 'stripPrefixRegex', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
