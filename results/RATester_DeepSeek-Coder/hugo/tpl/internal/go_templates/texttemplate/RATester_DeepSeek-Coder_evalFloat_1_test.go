package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalFloat_1(t *testing.T) {
	type testCase struct {
		name     string
		typ      reflect.Type
		node     parse.Node
		expected reflect.Value
	}

	testCases := []testCase{
		{
			name:     "valid float",
			typ:      reflect.TypeOf(float64(0)),
			node:     &parse.NumberNode{IsFloat: true, Float64: 123.456},
			expected: reflect.ValueOf(123.456),
		},
		{
			name:     "invalid float",
			typ:      reflect.TypeOf(float64(0)),
			node:     &parse.NumberNode{IsFloat: false, Float64: 123.456},
			expected: reflect.ValueOf(nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &state{}
			result := s.evalFloat(tc.typ, tc.node)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
