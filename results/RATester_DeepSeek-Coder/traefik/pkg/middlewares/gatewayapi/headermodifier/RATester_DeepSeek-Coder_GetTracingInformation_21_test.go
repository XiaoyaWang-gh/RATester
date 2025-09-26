package headermodifier

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rhm := &responseHeaderModifier{
		name: "test_modifier",
	}

	name, typeName, spanKind := rhm.GetTracingInformation()

	if name != "test_modifier" {
		t.Errorf("Expected name to be 'test_modifier', got %s", name)
	}

	if typeName != responseHeaderModifierTypeName {
		t.Errorf("Expected typeName to be '%s', got %s", responseHeaderModifierTypeName, typeName)
	}

	if spanKind != trace.SpanKindUnspecified {
		t.Errorf("Expected spanKind to be '%v', got %v", trace.SpanKindUnspecified, spanKind)
	}
}
