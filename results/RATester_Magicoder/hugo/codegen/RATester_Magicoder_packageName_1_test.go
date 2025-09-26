package codegen

import (
	"fmt"
	"go/ast"
	"testing"
)

func TestpackageName_1(t *testing.T) {
	type args struct {
		e ast.Expr
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := packageName(tt.args.e); got != tt.want {
				t.Errorf("packageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
