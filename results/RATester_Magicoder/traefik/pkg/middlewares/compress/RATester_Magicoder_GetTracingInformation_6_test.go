package compress

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	compress := &compress{
		name: "test",
	}
	name, typeName, kind := compress.GetTracingInformation()
	if name != "test" {
		t.Errorf("Expected name to be 'test', but got %s", name)
	}
	if typeName != "" {
		t.Errorf("Expected typeName to be '', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
