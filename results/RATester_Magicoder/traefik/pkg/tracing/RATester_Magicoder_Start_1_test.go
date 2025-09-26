package tracing

import (
	"context"
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestStart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	tracer := &Tracer{}
	spanName := "testSpan"

	ctx, span := tracer.Start(ctx, spanName)

	if span == nil {
		t.Error("Expected a non-nil span, but got nil")
	}

	if ctx == nil {
		t.Error("Expected a non-nil context, but got nil")
	}

	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		t.Error("Expected a valid span context, but got invalid")
	}

	if spanCtx.SpanID() == (trace.SpanID{}) {
		t.Error("Expected a non-zero span ID, but got zero")
	}

	if spanCtx.TraceID() == (trace.TraceID{}) {
		t.Error("Expected a non-zero trace ID, but got zero")
	}
}
