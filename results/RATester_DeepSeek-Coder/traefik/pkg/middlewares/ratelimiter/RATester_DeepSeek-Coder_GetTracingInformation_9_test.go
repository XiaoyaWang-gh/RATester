package ratelimiter

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rl := &rateLimiter{
		name:  "testRateLimiter",
		rate:  100,
		burst: 10,
		ttl:   60,
	}

	name, typeName, spanKind := rl.GetTracingInformation()

	if name != "testRateLimiter" {
		t.Errorf("Expected name to be 'testRateLimiter', got %s", name)
	}

	if typeName != "rateLimiter" {
		t.Errorf("Expected typeName to be 'rateLimiter', got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be trace.SpanKindInternal, got %v", spanKind)
	}
}
