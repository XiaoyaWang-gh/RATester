package fiber

import (
	"fmt"
	"testing"
)

func TestFindNextNonEscapedCharsetPosition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		search   string
		charset  []byte
		expected int
	}

	tests := []test{
		{"", []byte{'a'}, -1},
		{"a", []byte{'a'}, 0},
		{"aa", []byte{'a'}, 1},
		{"aaa", []byte{'a'}, 2},
		{"aaa\\a", []byte{'a'}, 5},
		{"aaa\\\\a", []byte{'a'}, 5},
		{"aaa\\\\\\a", []byte{'a'}, -1},
		{"aaa\\\\\\\\a", []byte{'a'}, 7},
		{"aaa\\\\\\\\\\a", []byte{'a'}, -1},
	}

	for _, test := range tests {
		result := findNextNonEscapedCharsetPosition(test.search, test.charset)
		if result != test.expected {
			t.Errorf("findNextNonEscapedCharsetPosition(%s, %s) = %d; want %d", test.search, string(test.charset), result, test.expected)
		}
	}
}
