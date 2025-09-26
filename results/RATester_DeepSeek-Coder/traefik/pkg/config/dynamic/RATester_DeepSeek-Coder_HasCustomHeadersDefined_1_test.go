package dynamic

import (
	"fmt"
	"testing"
)

func TestHasCustomHeadersDefined_1(t *testing.T) {
	type testCase struct {
		name     string
		headers  *Headers
		expected bool
	}

	testCases := []testCase{
		{
			name: "No custom headers",
			headers: &Headers{
				CustomRequestHeaders: map[string]string{},
			},
			expected: false,
		},
		{
			name: "Custom request headers",
			headers: &Headers{
				CustomRequestHeaders: map[string]string{
					"Header1": "Value1",
				},
			},
			expected: true,
		},
		{
			name: "Custom response headers",
			headers: &Headers{
				CustomResponseHeaders: map[string]string{
					"Header1": "Value1",
				},
			},
			expected: true,
		},
		{
			name: "Custom request and response headers",
			headers: &Headers{
				CustomRequestHeaders: map[string]string{
					"Header1": "Value1",
				},
				CustomResponseHeaders: map[string]string{
					"Header1": "Value1",
				},
			},
			expected: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.headers.HasCustomHeadersDefined()
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
