package hugocontext

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func TestParse_4(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output ast.Node
	}{
		{
			name:  "Test case 1",
			input: "{{<hugoContext>}}",
			output: &HugoContext{
				Closing: false,
			},
		},
		{
			name:  "Test case 2",
			input: "{{</hugoContext>}}",
			output: &HugoContext{
				Closing: true,
			},
		},
		{
			name:  "Test case 3",
			input: "{{<hugoContext attr=value>}}",
			output: &HugoContext{
				Closing: false,
			},
		},
	}

	parser := &hugoContextParser{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			reader := text.NewReader([]byte(test.input))
			node := parser.Parse(nil, reader, nil)

			if !reflect.DeepEqual(node, test.output) {
				t.Errorf("Expected %v, got %v", test.output, node)
			}
		})
	}
}
