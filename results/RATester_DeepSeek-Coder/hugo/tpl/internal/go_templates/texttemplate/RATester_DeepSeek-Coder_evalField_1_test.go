package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalField_1(t *testing.T) {
	type testStruct struct {
		Field string
	}

	testCases := []struct {
		name     string
		receiver reflect.Value
		field    string
		node     parse.Node
		args     []parse.Node
		final    reflect.Value
		expected reflect.Value
	}{
		{
			name:     "valid field",
			receiver: reflect.ValueOf(testStruct{Field: "test"}),
			field:    "Field",
			node:     &parse.FieldNode{},
			args:     []parse.Node{},
			final:    reflect.Value{},
			expected: reflect.ValueOf("test"),
		},
		{
			name:     "invalid field",
			receiver: reflect.ValueOf(testStruct{Field: "test"}),
			field:    "InvalidField",
			node:     &parse.FieldNode{},
			args:     []parse.Node{},
			final:    reflect.Value{},
			expected: reflect.Value{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &state{tmpl: &Template{}}
			result := s.evalField(reflect.Value{}, tc.field, tc.node, tc.args, tc.final, tc.receiver)
			if result.Interface() != tc.expected.Interface() {
				t.Errorf("Expected %v, but got %v", tc.expected.Interface(), result.Interface())
			}
		})
	}
}
