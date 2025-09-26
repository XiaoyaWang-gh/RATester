package gateway

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestRegisterFilterFuncs_1(t *testing.T) {
	type args struct {
		group       string
		kind        string
		builderFunc BuildFilterFunc
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRegisterFilterFuncs",
			args: args{
				group: "testGroup",
				kind:  "testKind",
				builderFunc: func(name, namespace string) (string, *dynamic.Middleware, error) {
					return "testName", &dynamic.Middleware{}, nil
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

			p := &Provider{}
			p.RegisterFilterFuncs(tt.args.group, tt.args.kind, tt.args.builderFunc)
			if p.groupKindFilterFuncs[tt.args.group][tt.args.kind] == nil {
				t.Errorf("RegisterFilterFuncs() = %v, want %v", p.groupKindFilterFuncs[tt.args.group][tt.args.kind], tt.args.builderFunc)
			}
		})
	}
}
