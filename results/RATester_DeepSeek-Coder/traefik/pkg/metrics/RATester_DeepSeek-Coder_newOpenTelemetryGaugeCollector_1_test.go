package metrics

import (
	"fmt"
	"testing"
)

func TestNewOpenTelemetryGaugeCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gauge := newOpenTelemetryGaugeCollector()

	if gauge == nil {
		t.Errorf("newOpenTelemetryGaugeCollector() returned nil")
	}

	if gauge.values == nil {
		t.Errorf("newOpenTelemetryGaugeCollector() returned a gauge with nil values")
	}
}
