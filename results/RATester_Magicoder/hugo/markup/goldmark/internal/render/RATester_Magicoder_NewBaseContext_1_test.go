package render

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/yuin/goldmark/ast"
)

func TestNewBaseContext_1(t *testing.T) {
	type args struct {
		rctx            *Context
		renderer        any
		n               ast.Node
		src             []byte
		getSourceSample func() []byte
		ordinal         int
	}
	tests := []struct {
		name string
		args args
		want hooks.BaseContext
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

			if got := NewBaseContext(tt.args.rctx, tt.args.renderer, tt.args.n, tt.args.src, tt.args.getSourceSample, tt.args.ordinal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
