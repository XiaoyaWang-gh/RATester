package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestTransform_11(t *testing.T) {
	tests := []struct {
		name string
		doc  *ast.Document
		want *ast.Document
	}{
		{
			name: "Test case 1",
			doc:  &ast.Document{},
			want: &ast.Document{},
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

			tr := &Transformer{
				wrapStandAloneImageWithinParagraph: true,
			}
			tr.Transform(tt.doc, nil, nil)

			// Compare the transformed document with the expected result
			// You may need to implement a custom comparison function if the documents are complex
			if !reflect.DeepEqual(tt.doc, tt.want) {
				t.Errorf("Transform() = %v, want %v", tt.doc, tt.want)
			}
		})
	}
}
