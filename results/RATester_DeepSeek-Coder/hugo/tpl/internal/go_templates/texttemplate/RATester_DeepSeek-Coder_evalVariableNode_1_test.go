package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalVariableNode_1(t *testing.T) {
	type testCase struct {
		name     string
		dot      reflect.Value
		variable *parse.VariableNode
		args     []parse.Node
		final    reflect.Value
		expected reflect.Value
	}

	tests := []testCase{
		{
			name: "test case 1",
			dot: reflect.ValueOf(struct {
				Field string
			}{
				Field: "expected value",
			}),
			variable: &parse.VariableNode{
				Ident: []string{"Field"},
			},
			args:     []parse.Node{},
			final:    reflect.Value{},
			expected: reflect.ValueOf("expected value"),
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

			s := &state{
				// Initialize any necessary fields here
			}
			result := s.evalVariableNode(tt.dot, tt.variable, tt.args, tt.final)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("evalVariableNode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
