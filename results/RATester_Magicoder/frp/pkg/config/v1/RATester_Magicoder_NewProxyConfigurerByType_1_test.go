package v1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProxyConfigurerByType_1(t *testing.T) {
	type args struct {
		proxyType ProxyType
	}
	tests := []struct {
		name string
		args args
		want ProxyConfigurer
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

			if got := NewProxyConfigurerByType(tt.args.proxyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxyConfigurerByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
