package fiber

import (
	"fmt"
	"testing"
)

func TestFindLastCharsetPosition_1(t *testing.T) {
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
		{"", []byte(""), -1},
		{"Hello, world", []byte("o"), 7},
		{"Hello, world", []byte("x"), -1},
		{"Hello, world", []byte("l"), 3},
		{"Hello, world", []byte("Hwxl"), 7},
		{"Hello, world", []byte("eH"), 1},
	}

	for _, test := range tests {
		if output := findLastCharsetPosition(test.search, test.charset); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, recieved: %v", test.search, test.expected, output)
		}
	}
}
