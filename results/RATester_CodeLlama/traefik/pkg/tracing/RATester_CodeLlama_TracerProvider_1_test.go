package tracing

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTracerProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Span{}
	assert.Equal(t, s.TracerProvider(), s.tracerProvider)
}
