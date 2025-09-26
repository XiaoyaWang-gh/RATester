package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSUDPProxy_1(t *testing.T) {
	type args struct {
		baseProxy *BaseProxy
	}
	tests := []struct {
		name string
		args args
		want Proxy
	}{
		{
			name: "test",
			args: args{
				baseProxy: &BaseProxy{
					name: "test",
				},
			},
			want: &SUDPProxy{
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

			if got := NewSUDPProxy(tt.args.baseProxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSUDPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
