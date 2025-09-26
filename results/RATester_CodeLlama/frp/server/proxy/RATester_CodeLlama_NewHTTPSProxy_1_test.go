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
		{
			name: "test1",
			args: args{
				baseProxy: &BaseProxy{
					name: "test1",
				},
			},
			want: &HTTPSProxy{
				BaseProxy: &BaseProxy{
					name: "test1",
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

			if got := NewHTTPSProxy(tt.args.baseProxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPSProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
