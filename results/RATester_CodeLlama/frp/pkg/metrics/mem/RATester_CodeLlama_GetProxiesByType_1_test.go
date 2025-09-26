package mem

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/metric"
	"github.com/zeebo/assert"
)

func TestGetProxiesByType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &serverMetrics{
		info: &ServerStatistics{
			ProxyTypeCounts: map[string]metric.Counter{
				"tcp":  metric.NewCounter(),
				"http": metric.NewCounter(),
			},
			ProxyStatistics: map[string]*ProxyStatistics{
				"proxy1": {
					ProxyType: "tcp",
				},
				"proxy2": {
					ProxyType: "http",
				},
				"proxy3": {
					ProxyType: "tcp",
				},
			},
		},
	}
	proxies := m.GetProxiesByType("tcp")
	assert.Equal(t, 2, len(proxies))
	assert.Equal(t, "proxy1", proxies[0].Name)
	assert.Equal(t, "proxy3", proxies[1].Name)
}
