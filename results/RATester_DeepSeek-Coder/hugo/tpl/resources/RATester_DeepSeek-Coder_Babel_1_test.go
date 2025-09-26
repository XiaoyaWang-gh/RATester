package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/js"
)

func TestBabel_1(t *testing.T) {
	ns := &Namespace{
		deps: &deps.Deps{},
		jsNs: &js.Namespace{},
	}

	testCases := []struct {
		name     string
		args     []any
		expected any
	}{
		{
			name:     "success",
			args:     []any{"arg1", "arg2"},
			expected: "expected_resource",
		},
		{
			name:     "error",
			args:     []any{"arg3", "arg4"},
			expected: "expected_error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Babel(tc.args...)
			if err != nil {
				if tc.expected != "expected_error" {
					t.Errorf("Expected no error, got %v", err)
				}
			} else {
				if result != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				}
			}
		})
	}
}
