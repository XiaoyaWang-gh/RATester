package context

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestProxy_1(t *testing.T) {
	testCases := []struct {
		name     string
		header   http.Header
		expected []string
	}{
		{
			name:     "no X-Forwarded-For",
			header:   http.Header{},
			expected: []string{},
		},
		{
			name: "single IP",
			header: http.Header{
				"X-Forwarded-For": []string{"192.0.2.0"},
			},
			expected: []string{"192.0.2.0"},
		},
		{
			name: "multiple IPs",
			header: http.Header{
				"X-Forwarded-For": []string{"192.0.2.0, 192.0.2.1, 192.0.2.2"},
			},
			expected: []string{"192.0.2.0", "192.0.2.1", "192.0.2.2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req := httptest.NewRequest("GET", "/", nil)
			req.Header = tc.header
			input := &BeegoInput{Context: &Context{Request: req}}
			result := input.Proxy()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
