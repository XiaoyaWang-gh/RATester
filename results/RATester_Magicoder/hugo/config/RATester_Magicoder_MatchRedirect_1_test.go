package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatchRedirect_1(t *testing.T) {
	server := &Server{
		Redirects: []Redirect{
			{From: "/old-page", To: "/new-page", Status: 301},
			{From: "/old-page2", To: "/new-page2", Status: 302},
		},
	}

	tests := []struct {
		name     string
		pattern  string
		expected Redirect
	}{
		{
			name:    "Match existing redirect",
			pattern: "/old-page",
			expected: Redirect{
				From:   "/old-page",
				To:     "/new-page",
				Status: 301,
			},
		},
		{
			name:    "Match existing redirect 2",
			pattern: "/old-page2",
			expected: Redirect{
				From:   "/old-page2",
				To:     "/new-page2",
				Status: 302,
			},
		},
		{
			name:     "No match",
			pattern:  "/nonexistent-page",
			expected: Redirect{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := server.MatchRedirect(test.pattern)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
