package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestexecFilter_1(t *testing.T) {
	type args struct {
		context *beecontext.Context
		urlPath string
		pos     int
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

			if got := tt.p.execFilter(tt.args.context, tt.args.urlPath, tt.args.pos); got != tt.want {
				t.Errorf("execFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
