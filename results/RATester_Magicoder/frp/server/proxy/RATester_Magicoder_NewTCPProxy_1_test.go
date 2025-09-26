package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewTCPProxy_1(t *testing.T) {
	type args struct {
		baseProxy *BaseProxy
	}
	tests := []struct {
		name string
		args args
		want Proxy
	}{
		{
			name: "TestNewTCPProxy",
			args: args{
				baseProxy: &BaseProxy{
					configurer: &v1.TCPProxyConfig{},
				},
			},
			want: &TCPProxy{
				BaseProxy: &BaseProxy{
					configurer: &v1.TCPProxyConfig{},
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

			if got := NewTCPProxy(tt.args.baseProxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTCPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
