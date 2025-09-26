package prometheus

import (
	"fmt"
	"testing"
)

func TestNewServerMetrics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	metrics := newServerMetrics()

	if metrics.clientCount == nil {
		t.Errorf("clientCount should not be nil")
	}

	if metrics.proxyCount == nil {
		t.Errorf("proxyCount should not be nil")
	}

	if metrics.connectionCount == nil {
		t.Errorf("connectionCount should not be nil")
	}

	if metrics.trafficIn == nil {
		t.Errorf("trafficIn should not be nil")
	}

	if metrics.trafficOut == nil {
		t.Errorf("trafficOut should not be nil")
	}
}
