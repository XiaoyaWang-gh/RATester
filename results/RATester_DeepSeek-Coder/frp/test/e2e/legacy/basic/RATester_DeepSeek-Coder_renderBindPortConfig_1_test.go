package basic

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/framework/consts"
)

func TestRenderBindPortConfig_1(t *testing.T) {
	type args struct {
		protocol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with kcp protocol",
			args: args{protocol: "kcp"},
			want: fmt.Sprintf(`kcp_bind_port = {{ .%s }}`, consts.PortServerName),
		},
		{
			name: "Test with quic protocol",
			args: args{protocol: "quic"},
			want: fmt.Sprintf(`quic_bind_port = {{ .%s }}`, consts.PortServerName),
		},
		{
			name: "Test with unknown protocol",
			args: args{protocol: "unknown"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := renderBindPortConfig(tt.args.protocol); got != tt.want {
				t.Errorf("renderBindPortConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
