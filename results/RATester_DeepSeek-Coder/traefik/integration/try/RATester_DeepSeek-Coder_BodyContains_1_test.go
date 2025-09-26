package try

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBodyContains_1(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		values   []string
		expected error
	}{
		{
			name:     "contains value",
			body:     "Hello, World",
			values:   []string{"World"},
			expected: nil,
		},
		{
			name:     "does not contain value",
			body:     "Hello, World",
			values:   []string{"Universe"},
			expected: errors.New("could not find 'Universe' in body 'Hello, World'"),
		},
		{
			name:     "empty body",
			body:     "",
			values:   []string{"Universe"},
			expected: errors.New("could not find 'Universe' in body ''"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			res := &http.Response{
				Body: io.NopCloser(strings.NewReader(tc.body)),
			}

			err := BodyContains(tc.values...)(res)
			if !errors.Is(err, tc.expected) {
				t.Errorf("Expected error '%v', got '%v'", tc.expected, err)
			}
		})
	}
}
