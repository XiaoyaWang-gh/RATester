package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequestHeader_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected string
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test Case 1",
			key:      "Content-Type",
			expected: "application/json",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
			},
		},
		{
			name:     "Test Case 2",
			key:      "Authorization",
			expected: "Bearer token",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Authorization": []string{"Bearer token"},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.requestHeader(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
