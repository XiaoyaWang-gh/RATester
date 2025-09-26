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
			name: "Test with nil hostResolverConfig",
			args: args{
				hostResolverConfig: nil,
			},
			want: &RequestDecorator{
				hostResolver: nil,
			},
		},
		{
			name: "Test with non-nil hostResolverConfig",
			args: args{
				hostResolverConfig: &types.HostResolverConfig{
					CnameFlattening: true,
					ResolvConfig:    "test",
					ResolvDepth:     1,
				},
			},
			want: &RequestDecorator{
				hostResolver: &Resolver{
					CnameFlattening: true,
					ResolvConfig:    "test",
					ResolvDepth:     1,
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

			got := New(tt.args.hostResolverConfig)
			if !reflect.DeepEqual(got.hostResolver, tt.want.hostResolver) {
				t.Errorf("New() = %v, want %v", got.hostResolver, tt.want.hostResolver)
			}
		})
	}
}
