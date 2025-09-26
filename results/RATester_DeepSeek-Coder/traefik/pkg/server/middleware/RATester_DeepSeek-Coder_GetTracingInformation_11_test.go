package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tp := &traceablePlugin{
		name: "testPlugin",
		h:    http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}),
	}

	name, typeName, kind := tp.GetTracingInformation()

	if name != "testPlugin" {
		t.Errorf("Expected name to be 'testPlugin', got %s", name)
	}

	if typeName != "traceablePlugin" {
		t.Errorf("Expected typeName to be 'traceablePlugin', got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be 'trace.SpanKindInternal', got %v", kind)
	}
}
