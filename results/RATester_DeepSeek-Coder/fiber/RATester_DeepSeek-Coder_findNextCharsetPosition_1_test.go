package fiber

import (
	"fmt"
	"testing"
)

func TestFindNextCharsetPosition_1(t *testing.T) {
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
		{"hello world", []byte("hw"), 0},
		{"hello world", []byte("z"), -1},
		{"hello world", []byte("l"), 2},
		{"hello world", []byte("o"), 4},
		{"hello world", []byte(" "), 5},
	}

	for _, test := range tests {
		if output := findNextCharsetPosition(test.search, test.charset); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, recieved: %v", test.search, test.expected, output)
		}
	}
}
