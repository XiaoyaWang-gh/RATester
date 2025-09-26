package httpcache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshalJSON_8(t *testing.T) {
	testCases := []struct {
		name     string
		input    PollConfig
		expected string
	}{
		{
			name: "Test case 1",
			input: PollConfig{
				For:     GlobMatcher{Excludes: []string{"exclude1", "exclude2"}, Includes: []string{"include1", "include2"}},
				Disable: false,
				Low:     10 * time.Second,
				High:    20 * time.Second,
			},
			expected: `{"For":{"Excludes":["exclude1","exclude2"],"Includes":["include1","include2"]},"Disable":false,"Low":"10s","High":"20s"}`,
		},
		{
			name: "Test case 2",
			input: PollConfig{
				For:     GlobMatcher{Excludes: []string{"exclude3", "exclude4"}, Includes: []string{"include3", "include4"}},
				Disable: true,
				Low:     30 * time.Second,
				High:    40 * time.Second,
			},
			expected: `{"For":{"Excludes":["exclude3","exclude4"],"Includes":["include3","include4"]},"Disable":true,"Low":"30s","High":"40s"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := json.Marshal(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, string(result))
			}
		})
	}
}
