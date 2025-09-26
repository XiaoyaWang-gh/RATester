package tracing

import (
	"context"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestTracerFromContext_1(t *testing.T) {
	testCases := []struct {
		desc     string
		ctx      context.Context
		expected *Tracer
	}{
		{
			desc:     "given a valid context with a valid tracer",
			ctx:      context.WithValue(context.Background(), "tracer", &Tracer{}),
			expected: &Tracer{},
		},
		{
			desc: "given a valid context with an invalid tracer",
			ctx:  context.WithValue(context.Background(), "tracer", "invalid"),
		},
		{
			desc: "given an invalid context",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tracer := TracerFromContext(test.ctx)
			assert.Equal(t, test.expected, tracer)
		})
	}
}
