package replacepathregex

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rp := &replacePathRegex{
		name: "test_name",
	}
	name, typeName, kind := rp.GetTracingInformation()

	if name != "test_name" {
		t.Errorf("Expected name to be 'test_name', but got %s", name)
	}

	if typeName != "replacePathRegex" {
		t.Errorf("Expected typeName to be 'replacePathRegex', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
