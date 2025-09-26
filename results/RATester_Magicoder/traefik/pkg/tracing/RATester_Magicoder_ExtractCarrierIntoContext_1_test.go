package tracing

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestExtractCarrierIntoContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	headers := make(http.Header)
	headers.Set("traceparent", "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")
	headers.Set("tracestate", "congo=t61rcWkgMzE")

	newCtx := ExtractCarrierIntoContext(ctx, headers)

	if newCtx.Value("traceparent") != "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01" {
		t.Errorf("Expected traceparent to be '00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01', but got %v", newCtx.Value("traceparent"))
	}

	if newCtx.Value("tracestate") != "congo=t61rcWkgMzE" {
		t.Errorf("Expected tracestate to be 'congo=t61rcWkgMzE', but got %v", newCtx.Value("tracestate"))
	}
}
