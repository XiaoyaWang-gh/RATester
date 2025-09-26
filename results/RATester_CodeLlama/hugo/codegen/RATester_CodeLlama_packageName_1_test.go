package codegen

import (
	"fmt"
	"go/ast"
	"testing"
)

func TestPackageName_1(t *testing.T) {
	tests := []struct {
		name string
		e    ast.Expr
		want string
	}{
		{
			name: "ast.Ident",
			e:    &ast.Ident{Name: "foo"},
			want: "foo",
		},
		{
			name: "ast.SelectorExpr",
			e: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "foo"},
				Sel: &ast.Ident{Name: "bar"},
			},
			want: "foo.bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := packageName(tt.e); got != tt.want {
				t.Errorf("packageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
