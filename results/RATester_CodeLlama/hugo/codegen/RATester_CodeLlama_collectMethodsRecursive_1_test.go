package codegen

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"
)

func TestCollectMethodsRecursive_1(t *testing.T) {
	type args struct {
		pkg string
		f   []*ast.Field
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test_case_1",
			args{
				pkg: "github.com/golang/example/hello",
				f: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "Hello",
							},
						},
						Type: &ast.Ident{
							Name: "string",
						},
					},
					{
						Names: []*ast.Ident{
							{
								Name: "World",
							},
						},
						Type: &ast.Ident{
							Name: "string",
						},
					},
				},
			},
			[]string{
				"string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := collectMethodsRecursive(tt.args.pkg, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectMethodsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
