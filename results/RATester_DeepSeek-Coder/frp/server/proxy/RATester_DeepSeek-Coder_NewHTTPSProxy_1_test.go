package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPSProxy_1(t *testing.T) {
	type args struct {
		baseProxy *BaseProxy
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

			if got := NewHTTPSProxy(tt.args.baseProxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPSProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
