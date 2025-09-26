package template

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEscapeText_1(t *testing.T) {
	e := &escaper{
		textNodeEdits: make(map[*parse.TextNode][]byte),
	}

	tests := []struct {
		name     string
		context  context
		textNode *parse.TextNode
		want     []byte
	}{
		{
			name: "escape text in text node",
			context: context{
				state: stateText,
			},
			textNode: &parse.TextNode{
				Text: []byte("<hello>"),
			},
			want: []byte("&lt;hello&gt;"),
		},
		{
			name: "escape text in text node with different state",
			context: context{
				state: 1,
			},
			textNode: &parse.TextNode{
				Text: []byte("<hello>"),
			},
			want: []byte("<hello>"),
		},
		{
			name: "escape text in text node with special characters",
			context: context{
				state: stateText,
			},
			textNode: &parse.TextNode{
				Text: []byte("<hello\nworld>"),
			},
			want: []byte("&lt;hello\nworld&gt;"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e.escapeText(tt.context, tt.textNode)
			if !bytes.Equal(tt.textNode.Text, tt.want) {
				t.Errorf("escapeText() = %v, want %v", tt.textNode.Text, tt.want)
			}
		})
	}
}
