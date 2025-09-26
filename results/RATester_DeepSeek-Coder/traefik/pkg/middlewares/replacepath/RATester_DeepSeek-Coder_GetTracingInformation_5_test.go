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
		next: nil,
		path: "/new/path",
		name: "replacePath",
	}

	name, typeName, kind := r.GetTracingInformation()

	if name != "replacePath" {
		t.Errorf("Expected name to be 'replacePath', got %s", name)
	}

	if typeName != "replacePath" {
		t.Errorf("Expected typeName to be 'replacePath', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, got %v", kind)
	}
}
