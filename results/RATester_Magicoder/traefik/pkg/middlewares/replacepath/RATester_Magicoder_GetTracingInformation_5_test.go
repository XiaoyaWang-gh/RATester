package replacepath

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &replacePath{
		name: "testName",
		path: "testPath",
	}
	name, typeName, kind := r.GetTracingInformation()
	if name != "testName" {
		t.Errorf("Expected name to be 'testName', but got '%s'", name)
	}
	if typeName != "replacePath" {
		t.Errorf("Expected typeName to be 'replacePath', but got '%s'", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got '%v'", kind)
	}
}
