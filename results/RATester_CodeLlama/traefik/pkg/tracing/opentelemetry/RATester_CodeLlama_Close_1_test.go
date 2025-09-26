package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	tp := sdktrace.NewTracerProvider()
	tc := &tpCloser{provider: tp}

	if err := tc.Close(); err != nil {
		t.Fatalf("tc.Close() failed: %v", err)
	}

	if err := tp.Shutdown(ctx); err != nil {
		t.Fatalf("tp.Shutdown() failed: %v", err)
	}
}
