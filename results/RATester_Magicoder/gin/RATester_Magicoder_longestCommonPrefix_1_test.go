package gin

import (
	"fmt"
	"testing"
)

func TestlongestCommonPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		a, b string
		want int
	}

	tests := []test{
		{"flower", "flow", 4},
		{"dog", "dogs", 2},
		{"cat", "car", 2},
		{"abc", "abc", 3},
		{"abc", "abcd", 3},
		{"abc", "ab", 2},
		{"abc", "", 0},
		{"", "abc", 0},
		{"", "", 0},
	}

	for _, tt := range tests {
		got := longestCommonPrefix(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("longestCommonPrefix(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
