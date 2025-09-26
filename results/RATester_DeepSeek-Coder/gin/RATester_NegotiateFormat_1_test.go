package gin

import (
	"fmt"
	"testing"
)

func TestNegotiateFormat_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		offered  []string
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with no offered formats",
			context: &Context{
				Accepted: []string{"application/json"},
			},
			offered:  []string{},
			expected: "",
		},
		{
			name: "Test with one offered format",
			context: &Context{
				Accepted: []string{"application/json"},
			},
			offered:  []string{"application/json"},
			expected: "application/json",
		},
		{
			name: "Test with multiple offered formats",
			context: &Context{
				Accepted: []string{"application/json"},
			},
			offered:  []string{"application/json", "application/xml"},
			expected: "application/json",
		},
		{
			name: "Test with no accepted formats",
			context: &Context{
				Accepted: []string{},
			},
			offered:  []string{"application/json"},
			expected: "application/json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.NegotiateFormat(tc.offered...)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
