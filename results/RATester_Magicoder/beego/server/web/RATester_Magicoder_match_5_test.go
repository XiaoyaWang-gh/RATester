package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func Testmatch_5(t *testing.T) {
	type args struct {
		treePattern    string
		wildcardValues []string
		ctx            *context.Context
	}
	tests := []struct {
		name string
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

			leaf := &leafInfo{}
			if got := leaf.match(tt.args.treePattern, tt.args.wildcardValues, tt.args.ctx); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
