package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestErrorf_3(t *testing.T) {
	testCases := []struct {
		name     string
		k        ErrorCode
		node     parse.Node
		line     int
		f        string
		args     []any
		expected *Error
	}{
		{
			name: "test case 1",
			k:    1,
			node: &parse.ListNode{},
			line: 10,
			f:    "error: %s",
			args: []any{"test"},
			expected: &Error{
				ErrorCode:   1,
				Node:        &parse.ListNode{},
				Name:        "",
				Line:        10,
				Description: "error: test",
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := errorf(tc.k, tc.node, tc.line, tc.f, tc.args...)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
