package pagination

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestPage_1(t *testing.T) {
	type testCase struct {
		name     string
		request  *http.Request
		expected int
	}

	testCases := []testCase{
		{
			name: "Test with valid page number",
			request: &http.Request{
				Form: url.Values{"p": []string{"5"}},
			},
			expected: 5,
		},
		{
			name: "Test with invalid page number",
			request: &http.Request{
				Form: url.Values{"p": []string{"0"}},
			},
			expected: 1,
		},
		{
			name: "Test with empty page number",
			request: &http.Request{
				Form: url.Values{"p": []string{""}},
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Paginator{
				Request: tc.request,
			}
			result := p.Page()
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
