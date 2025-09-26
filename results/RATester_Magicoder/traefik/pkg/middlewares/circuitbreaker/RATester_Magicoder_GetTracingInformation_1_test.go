package circuitbreaker

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cb := &circuitBreaker{
		name: "testCircuitBreaker",
	}
	name, typeName, kind := cb.GetTracingInformation()

	if name != "testCircuitBreaker" {
		t.Errorf("Expected name to be 'testCircuitBreaker', but got %s", name)
	}

	if typeName != "circuitBreaker" {
		t.Errorf("Expected typeName to be 'circuitBreaker', but got %s", typeName)
	}

	if kind != trace.SpanKindInternal {
		t.Errorf("Expected kind to be trace.SpanKindInternal, but got %v", kind)
	}
}
