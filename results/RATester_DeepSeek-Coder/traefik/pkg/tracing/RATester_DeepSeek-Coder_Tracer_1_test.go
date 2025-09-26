package tracing

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestTracer_1(t *testing.T) {
	tracer := &Tracer{}
	provider := TracerProvider{
		TracerProvider: noop.NewTracerProvider(),
		tracer:         tracer,
	}

	tests := []struct {
		name   string
		input  string
		output trace.Tracer
	}{
		{
			name:   "should return traefik tracer when name is traefik",
			input:  "github.com/traefik/traefik",
			output: tracer,
		},
		{
			name:   "should return default tracer when name is not traefik",
			input:  "some/other/name",
			output: noop.NewTracerProvider().Tracer("some/other/name"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := provider.Tracer(tt.input)
			if result != tt.output {
				t.Errorf("Expected %v, got %v", tt.output, result)
			}
		})
	}
}
