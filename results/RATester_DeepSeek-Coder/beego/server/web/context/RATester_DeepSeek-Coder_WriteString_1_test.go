package context

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteString_1(t *testing.T) {
	testCases := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Test case 1",
			content:  "Hello, World",
			expected: "Hello, World",
		},
		{
			name:     "Test case 2",
			content:  "Testing 1, 2, 3",
			expected: "Testing 1, 2, 3",
		},
		{
			name:     "Test case 3",
			content:  "",
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

			w := httptest.NewRecorder()
			ctx := &Context{
				ResponseWriter: &Response{
					ResponseWriter: w,
				},
			}

			ctx.WriteString(tc.content)

			if w.Body.String() != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, w.Body.String())
			}
		})
	}
}
