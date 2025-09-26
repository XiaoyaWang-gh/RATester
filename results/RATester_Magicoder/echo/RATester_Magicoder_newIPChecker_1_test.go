package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIPChecker_1(t *testing.T) {
	type args struct {
		configs []TrustOption
	}
	tests := []struct {
		name string
		args args
		want *ipChecker
	}{
		{
			name: "Test case 1",
			args: args{
				configs: []TrustOption{
					func(c *ipChecker) {
						c.trustLoopback = false
					},
					func(c *ipChecker) {
						c.trustLinkLocal = false
					},
					func(c *ipChecker) {
						c.trustPrivateNet = false
					},
				},
			},
			want: &ipChecker{
				trustLoopback:   false,
				trustLinkLocal:  false,
				trustPrivateNet: false,
			},
		},
		{
			name: "Test case 2",
			args: args{
				configs: []TrustOption{
					func(c *ipChecker) {
						c.trustLoopback = true
					},
					func(c *ipChecker) {
						c.trustLinkLocal = true
					},
					func(c *ipChecker) {
						c.trustPrivateNet = true
					},
				},
			},
			want: &ipChecker{
				trustLoopback:   true,
				trustLinkLocal:  true,
				trustPrivateNet: true,
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

			if got := newIPChecker(tt.args.configs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIPChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
