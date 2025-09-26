package attributes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestOpen_3(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *attributesBlock
		wantState parser.State
	}{
		{
			name:  "Test case 1",
			input: "[attr1=value1]",
			want: &attributesBlock{
				BaseBlock: ast.BaseBlock{},
			},
			wantState: parser.NoChildren,
		},
		{
			name:  "Test case 2",
			input: "[attr2=value2 attr3=value3]",
			want: &attributesBlock{
				BaseBlock: ast.BaseBlock{},
			},
			wantState: parser.NoChildren,
		},
		{
			name:      "Test case 3",
			input:     "invalid input",
			want:      nil,
			wantState: parser.RequireParagraph,
		},
	}

	parser := &attrParser{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			reader := text.NewReader([]byte(tt.input))
			got, gotState := parser.Open(nil, reader, nil)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Open() got = %v, want %v", got, tt.want)
			}
			if gotState != tt.wantState {
				t.Errorf("Open() gotState = %v, want %v", gotState, tt.wantState)
			}
		})
	}
}
