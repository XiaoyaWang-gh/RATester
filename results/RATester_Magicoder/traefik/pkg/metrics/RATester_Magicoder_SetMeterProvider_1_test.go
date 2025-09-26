package metrics

import (
	"context"
	"fmt"
	"testing"

	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func TestSetMeterProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()

	// Create a new MeterProvider
	mp := sdkmetric.NewMeterProvider()

	// Call the function under test
	SetMeterProvider(mp)

	// Verify that the global MeterProvider has been set
	if otel.GetMeterProvider() != mp {
		t.Errorf("Expected MeterProvider to be set")
	}

	// Force flush the MeterProvider
	if err := mp.ForceFlush(ctx); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Shutdown the MeterProvider
	if err := mp.Shutdown(ctx); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
