package mem

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/metric"
)

func TestNewProxy_6(t *testing.T) {
	type args struct {
		name      string
		proxyType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name:      "proxy1",
				proxyType: "type1",
			},
		},
		{
			name: "Test case 2",
			args: args{
				name:      "proxy2",
				proxyType: "type2",
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
			_, ok := m.info.ProxyTypeCounts[tt.args.proxyType]
			if !ok {
				t.Errorf("NewProxy() error, proxy type %s not found in ProxyTypeCounts", tt.args.proxyType)
			}
			_, ok = m.info.ProxyStatistics[tt.args.name]
			if !ok {
				t.Errorf("NewProxy() error, proxy name %s not found in ProxyStatistics", tt.args.name)
			}
		})
	}
}
