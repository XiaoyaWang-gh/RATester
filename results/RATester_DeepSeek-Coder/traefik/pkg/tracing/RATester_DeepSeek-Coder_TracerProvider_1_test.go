package tracing

import (
	"fmt"
	"testing"
)

func TestTracerProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tracerProvider := &TracerProvider{}
	span := Span{tracerProvider: tracerProvider}

	tp := span.TracerProvider()

	if tp != tracerProvider {
		t.Errorf("Expected TracerProvider to be %v, got %v", tracerProvider, tp)
	}
}
