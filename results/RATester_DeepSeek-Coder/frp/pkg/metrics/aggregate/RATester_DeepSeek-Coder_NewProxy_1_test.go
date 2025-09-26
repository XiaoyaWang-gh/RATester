package aggregate

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/server/metrics"
)

func TestNewProxy_1(t *testing.T) {
	type fields struct {
		ms []metrics.ServerMetrics
	}
	type args struct {
		name      string
		proxyType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
				ms: tt.fields.ms,
			}
			m.NewProxy(tt.args.name, tt.args.proxyType)
		})
	}
}
