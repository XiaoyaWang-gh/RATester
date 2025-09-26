package ratelimiter

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var rl *rateLimiter
	// act
	name, typeName, spanKind := rl.GetTracingInformation()
	// assert
	assert.Equal(t, "rateLimiter", name)
	assert.Equal(t, "rateLimiter", typeName)
	assert.Equal(t, trace.SpanKindInternal, spanKind)
}
