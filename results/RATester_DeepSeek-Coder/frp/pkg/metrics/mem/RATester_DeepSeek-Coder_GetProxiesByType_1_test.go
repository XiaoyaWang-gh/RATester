package mem

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetProxiesByType_1(t *testing.T) {
	type fields struct {
		info *ServerStatistics
		mu   sync.Mutex
	}
	type args struct {
		proxyType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*ProxyStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &serverMetrics{
				info: tt.fields.info,
				mu:   tt.fields.mu,
			}
			if got := m.GetProxiesByType(tt.args.proxyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serverMetrics.GetProxiesByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
