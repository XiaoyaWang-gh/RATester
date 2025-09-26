package vhost

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewHTTPReverseProxy_1(t *testing.T) {
	type args struct {
		option      HTTPReverseProxyOptions
		vhostRouter *Routers
	}
	tests := []struct {
		name string
		args args
		want *HTTPReverseProxy
	}{
		{
			name: "test1",
			args: args{
				option: HTTPReverseProxyOptions{
					ResponseHeaderTimeoutS: 10,
				},
				vhostRouter: &Routers{},
			},
			want: &HTTPReverseProxy{
				responseHeaderTimeout: time.Duration(10) * time.Second,
				vhostRouter:           &Routers{},
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

			if got := NewHTTPReverseProxy(tt.args.option, tt.args.vhostRouter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPReverseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
