package codegen

import (
	"fmt"
	"go/ast"
	"testing"
)

func TestPackageName_1(t *testing.T) {
	type args struct {
		e ast.Expr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Ident",
			args: args{
				e: &ast.Ident{
					Name: "test",
				},
			},
			want: "test",
		},
		{
			name: "Test SelectorExpr",
			args: args{
				e: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "fmt"},
					Sel: &ast.Ident{Name: "Println"},
				},
			},
			want: "fmt.Println",
		},
		{
			name: "Test nested SelectorExpr",
			args: args{
				e: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "fmt"},
						Sel: &ast.Ident{Name: "Println"},
					},
					Sel: &ast.Ident{Name: "test"},
				},
			},
			want: "fmt.Println.test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := packageName(tt.args.e); got != tt.want {
				t.Errorf("packageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
