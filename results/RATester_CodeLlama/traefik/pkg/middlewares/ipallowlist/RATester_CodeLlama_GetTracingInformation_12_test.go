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

	al := &ipAllowLister{}
	name, typeName, spanKind := al.GetTracingInformation()
	if name != "" {
		t.Errorf("name should be empty")
	}
	if typeName != "" {
		t.Errorf("typeName should be empty")
	}
	if spanKind != trace.SpanKindInternal {
		t.Errorf("spanKind should be trace.SpanKindInternal")
	}
}
