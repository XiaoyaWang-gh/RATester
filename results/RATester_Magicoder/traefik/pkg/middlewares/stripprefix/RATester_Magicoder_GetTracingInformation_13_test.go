package stripprefix

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sp := &stripPrefix{
		name: "testName",
	}
	name, typeName, kind := sp.GetTracingInformation()
	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}
	if typeName != "stripPrefix" {
		t.Errorf("Expected typeName to be 'stripPrefix', but got %s", typeName)
	}
	if kind != trace.SpanKindUnspecified {
		t.Errorf("Expected kind to be 'Unspecified', but got %s", kind)
	}
}
