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
			name: "Test case 1",
			args: args{v: true},
			want: func(c *ipChecker) {
				c.trustLoopback = true
			},
		},
		{
			name: "Test case 2",
			args: args{v: false},
			want: func(c *ipChecker) {
				c.trustLoopback = false
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

			if got := TrustLoopback(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrustLoopback() = %v, want %v", got, tt.want)
			}
		})
	}
}
