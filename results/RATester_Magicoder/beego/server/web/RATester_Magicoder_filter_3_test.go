package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func Testfilter_3(t *testing.T) {
	type args struct {
		ctx               *context.Context
		urlPath           string
		preFilterParams   map[string]string
		filterFunc        FilterFunc
		next              *FilterRouter
		tree              *Tree
		pattern           string
		returnOnOutput    bool
		resetParams       bool
		ValidRouterResult bool
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
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

			f := &FilterRouter{
				filterFunc:     tt.args.filterFunc,
				next:           tt.args.next,
				tree:           tt.args.tree,
				pattern:        tt.args.pattern,
				returnOnOutput: tt.args.returnOnOutput,
				resetParams:    tt.args.resetParams,
			}
			got, got1 := f.filter(tt.args.ctx, tt.args.urlPath, tt.args.preFilterParams)
			if got != tt.want {
				t.Errorf("filter() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("filter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
