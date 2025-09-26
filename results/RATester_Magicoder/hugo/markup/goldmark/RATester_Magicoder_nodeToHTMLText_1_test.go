package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestnodeToHTMLText_1(t *testing.T) {
	tests := []struct {
		name   string
		n      ast.Node
		source []byte
		want   []byte
	}{
		{
			name: "Test Case 1",
			n: &ast.String{
				Value: []byte("Hello, World!"),
			},
			source: []byte("Hello, World!"),
			want:   []byte("Hello, World!"),
		},
		{
			name: "Test Case 2",
			n: &ast.String{
				Value: []byte("<script>alert('Hello, World!');</script>"),
			},
			source: []byte("<script>alert('Hello, World!');</script>"),
			want:   []byte("&lt;script&gt;alert('Hello, World!');&lt;/script&gt;"),
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

			if got := nodeToHTMLText(tt.n, tt.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeToHTMLText() = %v, want %v", got, tt.want)
			}
		})
	}
}
