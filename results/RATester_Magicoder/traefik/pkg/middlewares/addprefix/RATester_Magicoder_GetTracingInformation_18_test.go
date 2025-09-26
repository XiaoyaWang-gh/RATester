package addprefix

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &addPrefix{
		name:   "testName",
		prefix: "testPrefix",
	}
	name, typeName, kind := a.GetTracingInformation()
	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got %s", name)
	}
	if typeName != "addPrefix" {
		t.Errorf("Expected typeName to be 'addPrefix', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %s", kind)
	}
}
