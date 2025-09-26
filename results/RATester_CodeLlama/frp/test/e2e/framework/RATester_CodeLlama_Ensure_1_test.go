package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestEnsure_1(t *testing.T) {
	t.Parallel()

	type args struct {
		fns []EnsureFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				fns: []EnsureFunc{
					func(r *request.Response) bool {
						return true
					},
				},
			},
		},
		{
			name: "case2",
			args: args{
				fns: []EnsureFunc{
					func(r *request.Response) bool {
						return false
					},
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

			e := &RequestExpect{
				f: &Framework{},
			}
			e.Ensure(tt.args.fns...)
		})
	}
}
