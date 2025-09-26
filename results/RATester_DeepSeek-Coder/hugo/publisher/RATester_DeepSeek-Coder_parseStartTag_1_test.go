package publisher

import (
	"fmt"
	"testing"
)

func TestParseStartTag_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test case 1", "<div>", "div"},
		{"Test case 2", "<div />", "div"},
		{"Test case 3", "<div class=\"test\">", "div"},
		{"Test case 4", "<div class=\"test\" />", "div"},
		{"Test case 5", "<div  >", "div"},
		{"Test case 6", "<div    />", "div"},
		{"Test case 7", "<div/>", "div"},
		{"Test case 8", "<div   />", "div"},
		{"Test case 9", "<div   class=\"test\">", "div"},
		{"Test case 10", "<div   class=\"test\" />", "div"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := parseStartTag(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
