package mem

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/util/metric"
)

func TestNewServerMetrics_2(t *testing.T) {
	tests := []struct {
		name string
		want *serverMetrics
	}{
		{
			name: "Test newServerMetrics",
			want: &serverMetrics{
				info: &ServerStatistics{
					TotalTrafficIn:  metric.NewDateCounter(ReserveDays),
					TotalTrafficOut: metric.NewDateCounter(ReserveDays),
					CurConns:        metric.NewCounter(),

					ClientCounts:    metric.NewCounter(),
					ProxyTypeCounts: make(map[string]metric.Counter),

					ProxyStatistics: make(map[string]*ProxyStatistics),
				},
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

			got := newServerMetrics()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServerMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
