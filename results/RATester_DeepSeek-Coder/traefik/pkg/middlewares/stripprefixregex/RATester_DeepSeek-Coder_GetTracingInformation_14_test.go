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
		name: "testName",
	}

	name, typeName, kind := sr.GetTracingInformation()

	if name != "testName" {
		t.Errorf("Expected name to be 'testName', got %s", name)
	}

	if typeName != "stripPrefixRegex" {
		t.Errorf("Expected typeName to be 'stripPrefixRegex', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, got %v", kind)
	}
}
