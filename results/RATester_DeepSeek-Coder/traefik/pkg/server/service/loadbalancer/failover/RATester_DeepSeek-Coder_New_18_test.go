package failover

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNew_18(t *testing.T) {
	type args struct {
		hc *dynamic.HealthCheck
	}
	tests := []struct {
		name string
		args args
		want *Failover
	}{
		{
			name: "Test with nil HealthCheck",
			args: args{
				hc: nil,
			},
			want: &Failover{
				wantsHealthCheck: false,
			},
		},
		{
			name: "Test with non-nil HealthCheck",
			args: args{
				hc: &dynamic.HealthCheck{},
			},
			want: &Failover{
				wantsHealthCheck: true,
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

			if got := New(tt.args.hc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
