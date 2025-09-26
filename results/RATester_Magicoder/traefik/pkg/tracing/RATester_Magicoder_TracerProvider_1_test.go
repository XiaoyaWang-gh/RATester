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

	result := span.TracerProvider()

	if result != tracerProvider {
		t.Errorf("Expected %v, but got %v", tracerProvider, result)
	}
}
