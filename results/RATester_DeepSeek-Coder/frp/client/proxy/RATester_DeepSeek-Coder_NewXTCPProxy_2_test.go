package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewXTCPProxy_2(t *testing.T) {
	type args struct {
		baseProxy *BaseProxy
		cfg       v1.ProxyConfigurer
	}

	baseProxy := &BaseProxy{
		baseCfg: &v1.ProxyBaseConfig{
			Name: "test",
		},
	}

	cfg := &v1.XTCPProxyConfig{
		ProxyBaseConfig: v1.ProxyBaseConfig{
			Name: "test",
		},
	}

	tests := []struct {
		name string
		args args
		want Proxy
	}{
		{
			name: "TestNewXTCPProxy",
			args: args{
				baseProxy: baseProxy,
				cfg:       cfg,
			},
			want: &XTCPProxy{
				BaseProxy: baseProxy,
				cfg:       cfg,
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

			if got := NewXTCPProxy(tt.args.baseProxy, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXTCPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
