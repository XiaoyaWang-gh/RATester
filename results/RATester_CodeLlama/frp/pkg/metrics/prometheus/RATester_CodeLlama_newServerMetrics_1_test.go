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

	m := newServerMetrics()
	if m.clientCount == nil {
		t.Error("clientCount is nil")
	}
	if m.proxyCount == nil {
		t.Error("proxyCount is nil")
	}
	if m.connectionCount == nil {
		t.Error("connectionCount is nil")
	}
	if m.trafficIn == nil {
		t.Error("trafficIn is nil")
	}
	if m.trafficOut == nil {
		t.Error("trafficOut is nil")
	}
}
