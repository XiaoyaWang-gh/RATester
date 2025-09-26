package mem

import (
	"fmt"
	"testing"
)

func TestNewServerMetrics_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	metrics := newServerMetrics()

	if metrics == nil {
		t.Error("newServerMetrics() should not return nil")
	}

	if metrics.info == nil {
		t.Error("newServerMetrics().info should not return nil")
	}

	if metrics.info.TotalTrafficIn == nil {
		t.Error("newServerMetrics().info.TotalTrafficIn should not return nil")
	}

	if metrics.info.TotalTrafficOut == nil {
		t.Error("newServerMetrics().info.TotalTrafficOut should not return nil")
	}

	if metrics.info.CurConns == nil {
		t.Error("newServerMetrics().info.CurConns should not return nil")
	}

	if metrics.info.ClientCounts == nil {
		t.Error("newServerMetrics().info.ClientCounts should not return nil")
	}

	if metrics.info.ProxyTypeCounts == nil {
		t.Error("newServerMetrics().info.ProxyTypeCounts should not return nil")
	}

	if metrics.info.ProxyStatistics == nil {
		t.Error("newServerMetrics().info.ProxyStatistics should not return nil")
	}
}
