package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestFilter_4(t *testing.T) {
	type args struct {
		ctx *beecontext.Context
	}
	tests := []struct {
		name string
		l    *logFilter
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

			if got := tt.l.Filter(tt.args.ctx); got != tt.want {
				t.Errorf("logFilter.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
