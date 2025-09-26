package util

import (
	"errors"
	"fmt"
	"testing"
)

func TestGenerateResponseErrorString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var testCases = []struct {
		summary  string
		err      error
		detailed bool
		expected string
	}{
		{
			summary:  "test summary",
			err:      errors.New("test error"),
			detailed: true,
			expected: "test error",
		},
		{
			summary:  "test summary",
			err:      errors.New("test error"),
			detailed: false,
			expected: "test summary",
		},
	}

	for _, tc := range testCases {
		actual := GenerateResponseErrorString(tc.summary, tc.err, tc.detailed)
		if actual != tc.expected {
			t.Errorf("expected %s, actual %s", tc.expected, actual)
		}
	}
}
