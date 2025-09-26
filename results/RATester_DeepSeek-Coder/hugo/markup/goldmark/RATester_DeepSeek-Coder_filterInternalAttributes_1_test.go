package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestFilterInternalAttributes_1(t *testing.T) {
	tests := []struct {
		name  string
		attrs []ast.Attribute
		want  []ast.Attribute
	}{
		{
			name: "Test with no internal attributes",
			attrs: []ast.Attribute{
				{Name: []byte("class"), Value: "test"},
				{Name: []byte("id"), Value: "test"},
			},
			want: []ast.Attribute{
				{Name: []byte("class"), Value: "test"},
				{Name: []byte("id"), Value: "test"},
			},
		},
		{
			name: "Test with internal attributes",
			attrs: []ast.Attribute{
				{Name: []byte("class"), Value: "test"},
				{Name: []byte("id"), Value: "test"},
				{Name: []byte("internal-class"), Value: "test"},
			},
			want: []ast.Attribute{
				{Name: []byte("class"), Value: "test"},
				{Name: []byte("id"), Value: "test"},
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

			r := &hookedRenderer{}
			got := r.filterInternalAttributes(tt.attrs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterInternalAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
