package testhelpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithRouters_1(t *testing.T) {
	type args struct {
		opts []func(*dynamic.Router) string
	}
	tests := []struct {
		name string
		args args
		want func(*dynamic.HTTPConfiguration)
	}{
		{
			name: "Test case 1",
			args: args{
				opts: []func(*dynamic.Router) string{
					func(r *dynamic.Router) string {
						r.EntryPoints = []string{"web"}
						return "router1"
					},
					func(r *dynamic.Router) string {
						r.EntryPoints = []string{"websecure"}
						return "router2"
					},
				},
			},
			want: func(c *dynamic.HTTPConfiguration) {
				c.Routers = map[string]*dynamic.Router{
					"router1": {
						EntryPoints: []string{"web"},
					},
					"router2": {
						EntryPoints: []string{"websecure"},
					},
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

			got := WithRouters(tt.args.opts...)
			gotCfg := &dynamic.HTTPConfiguration{}
			got(gotCfg)
			wantCfg := &dynamic.HTTPConfiguration{}
			tt.want(wantCfg)
			if !reflect.DeepEqual(gotCfg, wantCfg) {
				t.Errorf("WithRouters() = %v, want %v", gotCfg, wantCfg)
			}
		})
	}
}
