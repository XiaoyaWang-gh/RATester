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

	// Arrange
	// Act
	result := newOpenTelemetryGaugeCollector()
	// Assert
	if result == nil {
		t.Errorf("newOpenTelemetryGaugeCollector() = nil")
	}
}
