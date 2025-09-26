package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestRegisterProxyFactory_2(t *testing.T) {
	type args struct {
		proxyConfType reflect.Type
		factory       func(*BaseProxy, v1.ProxyConfigurer) Proxy
	}
	tests := []struct {
		name string
		args args
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

			RegisterProxyFactory(tt.args.proxyConfType, tt.args.factory)
		})
	}
}
