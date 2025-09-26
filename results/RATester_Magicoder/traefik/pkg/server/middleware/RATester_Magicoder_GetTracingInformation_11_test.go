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
		name: "test_plugin",
		h:    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}

	name, typeName, kind := tp.GetTracingInformation()

	if name != "test_plugin" {
		t.Errorf("Expected name to be 'test_plugin', but got %s", name)
	}

	if typeName != "http.Handler" {
		t.Errorf("Expected typeName to be 'http.Handler', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be 'trace.SpanKindInternal', but got %s", kind)
	}
}
