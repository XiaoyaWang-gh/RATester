package attributes

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/yuin/goldmark/ast"
)

func TestRenderASTAttributes_1(t *testing.T) {
	type args struct {
		w          hugio.FlexiWriter
		attributes []ast.Attribute
	}
	tests := []struct {
		name string
		args args
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

			RenderASTAttributes(tt.args.w, tt.args.attributes...)
		})
	}
}
