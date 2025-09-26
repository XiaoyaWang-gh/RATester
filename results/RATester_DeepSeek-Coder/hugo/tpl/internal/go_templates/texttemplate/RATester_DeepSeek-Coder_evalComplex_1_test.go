package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalComplex_1(t *testing.T) {
	type testCase struct {
		name     string
		typ      reflect.Type
		node     parse.Node
		expected reflect.Value
	}

	testCases := []testCase{
		{
			name:     "valid complex number",
			typ:      reflect.TypeOf(complex(1, 1)),
			node:     &parse.NumberNode{IsComplex: true, Complex128: complex(1, 1)},
			expected: reflect.ValueOf(complex(1, 1)),
		},
		{
			name:     "invalid complex number",
			typ:      reflect.TypeOf(complex(1, 1)),
			node:     &parse.NumberNode{IsComplex: false, Complex128: complex(1, 1)},
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
			result := s.evalComplex(tc.typ, tc.node)
			if result.Interface() != tc.expected.Interface() {
				t.Errorf("expected %v, got %v", tc.expected.Interface(), result.Interface())
			}
		})
	}
}
