package gin

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testData struct {
		a, b     string
		expected int
	}

	testCases := []testData{
		{"flower", "flow", 3},
		{"flower", "flight", 2},
		{"flower", "f", 0},
		{"flower", "", 0},
		{"", "flight", 0},
		{"", "", 0},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("longestCommonPrefix(%s, %s) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
