package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEscapeListConditionally_1(t *testing.T) {
	e := &escaper{
		ns: &nameSpace{
			set: make(map[string]*Template),
		},
		output: make(map[string]context),
	}

	testCases := []struct {
		name     string
		e        *escaper
		n        *parse.ListNode
		filter   func(*escaper, context) bool
		expected bool
	}{
		{
			name: "filter returns true",
			e:    e,
			n:    &parse.ListNode{},
			filter: func(e *escaper, c context) bool {
				return true
			},
			expected: true,
		},
		{
			name: "filter returns false",
			e:    e,
			n:    &parse.ListNode{},
			filter: func(e *escaper, c context) bool {
				return false
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, actual := tc.e.escapeListConditionally(context{}, tc.n, tc.filter)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
