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
		t.Error("Expected a non-nil gauge, but got nil")
	}

	if gauge.values == nil {
		t.Error("Expected a non-nil values map, but got nil")
	}
}
