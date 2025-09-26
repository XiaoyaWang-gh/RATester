package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestContentType_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with Content-Type header",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
				},
			},
			expected: "application/json",
		},
		{
			name: "Test with no Content-Type header",
			context: &Context{
				Request: &http.Request{
					Header: http.Header{},
				},
			},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.ContentType()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
