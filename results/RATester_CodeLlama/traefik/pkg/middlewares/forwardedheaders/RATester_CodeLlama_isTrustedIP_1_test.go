package forwardedheaders

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/ip"
)

func TestIsTrustedIP_1(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		x    *XForwarded
		args args
		want bool
	}{
		{
			name: "test1",
			x:    &XForwarded{ipChecker: &ip.Checker{}},
			args: args{ip: "127.0.0.1"},
			want: true,
		},
		{
			name: "test2",
			x:    &XForwarded{ipChecker: &ip.Checker{}},
			args: args{ip: "127.0.0.2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.x.isTrustedIP(tt.args.ip); got != tt.want {
				t.Errorf("isTrustedIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
