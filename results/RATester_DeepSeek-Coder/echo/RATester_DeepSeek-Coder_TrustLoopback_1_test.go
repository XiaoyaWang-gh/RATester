package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrustLoopback_1(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want TrustOption
	}{
		{
			name: "TestTrustLoopback_True",
			args: args{v: true},
			want: TrustOption(func(c *ipChecker) {
				c.trustLoopback = true
			}),
		},
		{
			name: "TestTrustLoopback_False",
			args: args{v: false},
			want: TrustOption(func(c *ipChecker) {
				c.trustLoopback = false
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := TrustLoopback(tt.args.v)
			if reflect.ValueOf(got) != reflect.ValueOf(tt.want) {
				t.Errorf("TrustLoopback() = %v, want %v", got, tt.want)
			}
		})
	}
}
