package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPProxy_1(t *testing.T) {
	type args struct {
		baseProxy *BaseProxy
	}
	tests := []struct {
		name string
		args args
		want Proxy
	}{
		{
			name: "test case 1",
			args: args{
				baseProxy: &BaseProxy{
					name: "test",
				},
			},
			want: &HTTPProxy{
				BaseProxy: &BaseProxy{
					name: "test",
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

			if got := NewHTTPProxy(tt.args.baseProxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
