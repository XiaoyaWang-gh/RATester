package mem

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/metric"
)

func TestCloseProxy_6(t *testing.T) {
	type args struct {
		name      string
		proxyType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCloseProxy",
			args: args{
				name:      "testProxy",
				proxyType: "testType",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &serverMetrics{
				info: &ServerStatistics{
					ProxyTypeCounts: make(map[string]metric.Counter),
					ProxyStatistics: make(map[string]*ProxyStatistics),
				},
			}
			m.NewProxy(tt.args.name, tt.args.proxyType)
			m.CloseProxy(tt.args.name, tt.args.proxyType)
			if _, ok := m.info.ProxyTypeCounts[tt.args.proxyType]; ok {
				t.Errorf("ProxyTypeCounts should be decreased")
			}
			if _, ok := m.info.ProxyStatistics[tt.args.name]; !ok {
				t.Errorf("ProxyStatistics should be updated")
			}
		})
	}
}
