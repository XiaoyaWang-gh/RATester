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
		rate:  10,
		burst: 100,
		ttl:   60,
	}

	name, typeName, spanKind := rl.GetTracingInformation()

	if name != "testRateLimiter" {
		t.Errorf("Expected name to be 'testRateLimiter', but got %s", name)
	}

	if typeName != "rateLimiter" {
		t.Errorf("Expected typeName to be 'rateLimiter', but got %s", typeName)
	}

	if spanKind != trace.SpanKindInternal {
		t.Errorf("Expected spanKind to be trace.SpanKindInternal, but got %v", spanKind)
	}
}
