package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewGeneralTCPProxy_1(t *testing.T) {
	type args struct {
		baseProxy       *BaseProxy
		proxyConfigurer v1.ProxyConfigurer
	}
	tests := []struct {
		name string
		args args
		want Proxy
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

			if got := NewGeneralTCPProxy(tt.args.baseProxy, tt.args.proxyConfigurer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeneralTCPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
