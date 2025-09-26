package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestMatch_5(t *testing.T) {
	type args struct {
		treePattern    string
		wildcardValues []string
		ctx            *context.Context
	}
	tests := []struct {
		name   string
		li     *leafInfo
		args   args
		wantOk bool
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

			leaf := &leafInfo{
				wildcards: tt.li.wildcards,
				regexps:   tt.li.regexps,
			}
			if gotOk := leaf.match(tt.args.treePattern, tt.args.wildcardValues, tt.args.ctx); gotOk != tt.wantOk {
				t.Errorf("leafInfo.match() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
