package tracing

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func TestInjectContextIntoCarrier_1(t *testing.T) {
	t.Run("TestInjectContextIntoCarrier", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// Create a new request
		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new context
		ctx := context.Background()

		// Inject the context into the request
		InjectContextIntoCarrier(req.WithContext(ctx))

		// Check if the context is correctly injected
		propagator := otel.GetTextMapPropagator()
		carrier := propagation.HeaderCarrier(req.Header)
		ctx = context.Background()
		ctx = propagator.Extract(ctx, carrier)
		if ctx == nil {
			t.Fatal("Context not injected correctly")
		}
	})
}
