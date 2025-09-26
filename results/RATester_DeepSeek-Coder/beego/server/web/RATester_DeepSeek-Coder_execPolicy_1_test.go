package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestExecPolicy_1(t *testing.T) {
	type args struct {
		cont    *context.Context
		urlPath string
	}
	tests := []struct {
		name string
		p    *ControllerRegister
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.execPolicy(tt.args.cont, tt.args.urlPath); got != tt.want {
				t.Errorf("execPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
