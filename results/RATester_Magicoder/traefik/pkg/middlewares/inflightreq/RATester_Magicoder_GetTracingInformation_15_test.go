package inflightreq

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &inFlightReq{
		name: "test",
	}
	name, typeName, kind := i.GetTracingInformation()
	if name != "test" {
		t.Errorf("Expected name to be 'test', but got %s", name)
	}
	if typeName != "inFlightReq" {
		t.Errorf("Expected typeName to be 'inFlightReq', but got %s", typeName)
	}
	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
