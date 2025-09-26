package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalFieldChain_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		dot      reflect.Value
		receiver reflect.Value
		node     parse.Node
		ident    []string
		args     []parse.Node
		final    reflect.Value
		expected reflect.Value
	}{
		{
			name:     "test case 1",
			dot:      reflect.ValueOf(nil),
			receiver: reflect.ValueOf(nil),
			node:     &parse.IdentifierNode{},
			ident:    []string{"field1", "field2"},
			args:     []parse.Node{&parse.IdentifierNode{}},
			final:    reflect.ValueOf(nil),
			expected: reflect.ValueOf(nil),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			s := &state{}
			result := s.evalFieldChain(tc.dot, tc.receiver, tc.node, tc.ident, tc.args, tc.final)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
