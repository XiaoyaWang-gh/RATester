package mem

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/metric"
)

func TestOpenConnection_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	metrics := &serverMetrics{
		info: &ServerStatistics{
			CurConns:        metric.NewCounter(),
			ProxyStatistics: make(map[string]*ProxyStatistics),
		},
	}

	name := "testProxy"
	metrics.OpenConnection(name, "")

	if metrics.info.CurConns.Count() != 1 {
		t.Errorf("Expected CurConns to be 1, but got %d", metrics.info.CurConns.Count())
	}

	proxyStats, ok := metrics.info.ProxyStatistics[name]
	if !ok {
		t.Errorf("Expected ProxyStatistics to contain %s, but it does not", name)
	}

	if proxyStats.CurConns.Count() != 1 {
		t.Errorf("Expected CurConns in ProxyStatistics to be 1, but got %d", proxyStats.CurConns.Count())
	}
}
