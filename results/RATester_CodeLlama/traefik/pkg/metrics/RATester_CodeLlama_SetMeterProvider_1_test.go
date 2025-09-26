package metrics

import (
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

	meterProvider := &sdkmetric.MeterProvider{}
	SetMeterProvider(meterProvider)
	if openTelemetryMeterProvider != meterProvider {
		t.Errorf("SetMeterProvider() = %v, want %v", openTelemetryMeterProvider, meterProvider)
	}
}
