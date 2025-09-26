package codegen

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"
)

func TestcollectMethodsRecursive_1(t *testing.T) {
	tests := []struct {
		name string
		pkg  string
		f    []*ast.Field
		want []string
	}{
		{
			name: "Test case 1",
			pkg:  "testpkg",
			f: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "Method1",
						},
					},
				},
				{
					Type: &ast.Ident{
						Obj: &ast.Object{
							Decl: &ast.TypeSpec{
								Type: &ast.InterfaceType{
									Methods: &ast.FieldList{
										List: []*ast.Field{
											{
												Names: []*ast.Ident{
													{
														Name: "Method2",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: []string{"Method1", "testpkg.Method2"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := collectMethodsRecursive(tt.pkg, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectMethodsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
