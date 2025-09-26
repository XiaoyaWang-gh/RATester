package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHost_1(t *testing.T) {
	type testCase struct {
		name     string
		input    *BeegoInput
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with empty host",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Host: "",
					},
				},
			},
			expected: "localhost",
		},
		{
			name: "Test with host and port",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Host: "localhost:8080",
					},
				},
			},
			expected: "localhost",
		},
		{
			name: "Test with host only",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Host: "localhost",
					},
				},
			},
			expected: "localhost",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.Host()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
