package vhost

import (
	"fmt"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
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
