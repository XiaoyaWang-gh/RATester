package dynamic

import (
	"fmt"
	"testing"
)

func TestHasSecureHeadersDefined_1(t *testing.T) {
	type testCase struct {
		name     string
		headers  *Headers
		expected bool
	}

	testCases := []testCase{
		{
			name: "empty headers",
			headers: &Headers{
				AllowedHosts: []string{},
			},
			expected: false,
		},
		{
			name: "has allowed hosts",
			headers: &Headers{
				AllowedHosts: []string{"example.com"},
			},
			expected: true,
		},
		// add more test cases here
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.headers.HasSecureHeadersDefined()
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
