package health

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewMonitor_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		cfg            v1.HealthCheckConfig
		addr           string
		statusNormalFn func()
		statusFailedFn func()
	}
	tests := []struct {
		name string
		args args
		want *Monitor
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

			got := NewMonitor(tt.args.ctx, tt.args.cfg, tt.args.addr, tt.args.statusNormalFn, tt.args.statusFailedFn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}
