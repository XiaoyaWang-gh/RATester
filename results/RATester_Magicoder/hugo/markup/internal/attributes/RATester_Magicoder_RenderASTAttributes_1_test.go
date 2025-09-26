package attributes

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestRenderASTAttributes_1(t *testing.T) {
	type args struct {
		w          *strings.Builder
		attributes []ast.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				w: &strings.Builder{},
				attributes: []ast.Attribute{
					{
						Name:  []byte("id"),
						Value: "test-id",
					},
					{
						Name:  []byte("class"),
						Value: "test-class",
					},
				},
			},
			want: " id=\"test-id\" class=\"test-class\"",
		},
		{
			name: "Test case 2",
			args: args{
				w: &strings.Builder{},
				attributes: []ast.Attribute{
					{
						Name:  []byte("onclick"),
						Value: "test-onclick",
					},
					{
						Name:  []byte("class"),
						Value: "test-class",
					},
				},
			},
			want: " class=\"test-class\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			RenderASTAttributes(tt.args.w, tt.args.attributes...)
			if got := tt.args.w.String(); got != tt.want {
				t.Errorf("RenderASTAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
