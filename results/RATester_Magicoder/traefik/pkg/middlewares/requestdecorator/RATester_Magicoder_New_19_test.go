package requestdecorator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNew_19(t *testing.T) {
	type args struct {
		hostResolverConfig *types.HostResolverConfig
	}
	tests := []struct {
		name string
		args args
		want *RequestDecorator
	}{
		{
			name: "Test case 1",
			args: args{
				hostResolverConfig: &types.HostResolverConfig{
					CnameFlattening: true,
					ResolvConfig:    "/etc/resolv.conf",
					ResolvDepth:     5,
				},
			},
			want: &RequestDecorator{
				hostResolver: &Resolver{
					CnameFlattening: true,
					ResolvConfig:    "/etc/resolv.conf",
					ResolvDepth:     5,
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				hostResolverConfig: &types.HostResolverConfig{
					CnameFlattening: false,
					ResolvConfig:    "/etc/resolv.conf",
					ResolvDepth:     10,
				},
			},
			want: &RequestDecorator{
				hostResolver: &Resolver{
					CnameFlattening: false,
					ResolvConfig:    "/etc/resolv.conf",
					ResolvDepth:     10,
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

			if got := New(tt.args.hostResolverConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
