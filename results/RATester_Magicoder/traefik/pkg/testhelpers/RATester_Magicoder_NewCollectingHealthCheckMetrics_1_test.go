package testhelpers

import (
	"fmt"
	"testing"
)

func TestNewCollectingHealthCheckMetrics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	metrics := NewCollectingHealthCheckMetrics()

	if metrics == nil {
		t.Error("Expected metrics to be not nil")
	}

	if metrics.Gauge == nil {
		t.Error("Expected Gauge to be not nil")
	}
}
