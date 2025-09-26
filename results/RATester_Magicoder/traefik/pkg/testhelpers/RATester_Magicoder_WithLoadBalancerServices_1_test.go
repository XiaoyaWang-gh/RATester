package testhelpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithLoadBalancerServices_1(t *testing.T) {
	type args struct {
		opts []func(service *dynamic.ServersLoadBalancer) string
	}
	tests := []struct {
		name string
		args args
		want func(*dynamic.HTTPConfiguration)
	}{
		{
			name: "Test case 1",
			args: args{
				opts: []func(service *dynamic.ServersLoadBalancer) string{
					func(service *dynamic.ServersLoadBalancer) string {
						return "service1"
					},
				},
			},
			want: func(c *dynamic.HTTPConfiguration) {
				c.Services = make(map[string]*dynamic.Service)
				b := &dynamic.ServersLoadBalancer{}
				name := "service1"
				c.Services[name] = &dynamic.Service{
					LoadBalancer: b,
				}
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := WithLoadBalancerServices(tt.args.opts...)
			httpConfig := &dynamic.HTTPConfiguration{}
			got(httpConfig)
			wantConfig := &dynamic.HTTPConfiguration{}
			tt.want(wantConfig)
			if !reflect.DeepEqual(httpConfig, wantConfig) {
				t.Errorf("WithLoadBalancerServices() = %v, want %v", httpConfig, wantConfig)
			}
		})
	}
}
