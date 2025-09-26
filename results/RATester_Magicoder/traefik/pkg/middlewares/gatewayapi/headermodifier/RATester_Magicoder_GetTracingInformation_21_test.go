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

	modifier := &responseHeaderModifier{
		name: "testModifier",
	}

	name, typeName, spanKind := modifier.GetTracingInformation()

	if name != "testModifier" {
		t.Errorf("Expected name to be 'testModifier', but got '%s'", name)
	}

	if typeName != responseHeaderModifierTypeName {
		t.Errorf("Expected typeName to be '%s', but got '%s'", responseHeaderModifierTypeName, typeName)
	}

	if spanKind != trace.SpanKindUnspecified {
		t.Errorf("Expected spanKind to be '%s', but got '%s'", trace.SpanKindUnspecified, spanKind)
	}
}
