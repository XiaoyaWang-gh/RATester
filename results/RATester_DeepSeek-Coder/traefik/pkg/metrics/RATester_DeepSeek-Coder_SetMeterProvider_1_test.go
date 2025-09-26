package metrics

import (
	"context"
	"fmt"
	"testing"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func TestSetMeterProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()

	testMeterProvider := &sdkmetric.MeterProvider{}
	SetMeterProvider(testMeterProvider)

	if openTelemetryMeterProvider != testMeterProvider {
		t.Errorf("Expected openTelemetryMeterProvider to be %v, but got %v", testMeterProvider, openTelemetryMeterProvider)
	}

	err := openTelemetryMeterProvider.ForceFlush(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	err = openTelemetryMeterProvider.Shutdown(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
