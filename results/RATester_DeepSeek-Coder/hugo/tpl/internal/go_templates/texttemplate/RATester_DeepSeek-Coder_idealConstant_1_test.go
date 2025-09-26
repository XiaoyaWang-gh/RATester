package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestIdealConstant_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		constant *parse.NumberNode
		expected reflect.Value
	}{
		{
			name: "Test with complex number",
			constant: &parse.NumberNode{
				IsComplex:  true,
				Complex128: 1 + 1i,
			},
			expected: reflect.ValueOf(1 + 1i),
		},
		{
			name: "Test with float number",
			constant: &parse.NumberNode{
				IsFloat: true,
				Float64: 1.1,
			},
			expected: reflect.ValueOf(1.1),
		},
		{
			name: "Test with int number",
			constant: &parse.NumberNode{
				IsInt: true,
				Int64: 1,
			},
			expected: reflect.ValueOf(1),
		},
		{
			name: "Test with uint number",
			constant: &parse.NumberNode{
				IsUint: true,
				Uint64: 1,
			},
			expected: reflect.ValueOf(uint(1)),
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
			result := s.idealConstant(tc.constant)
			if result.Interface() != tc.expected.Interface() {
				t.Errorf("Expected %v, but got %v", tc.expected.Interface(), result.Interface())
			}
		})
	}
}
