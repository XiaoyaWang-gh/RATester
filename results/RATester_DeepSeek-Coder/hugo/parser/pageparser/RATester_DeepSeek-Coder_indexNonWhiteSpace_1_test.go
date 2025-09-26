package pageparser

import (
	"fmt"
	"testing"
)

func TestIndexNonWhiteSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    []byte
		in       rune
		expected int
	}

	tests := []test{
		{[]byte("Hello, World!"), 'W', 7},
		{[]byte("Hello, World!"), 'z', -1},
		{[]byte("Hello,  World!"), ' ', 7},
		{[]byte("Hello, World!"), 'H', 0},
	}

	for _, test := range tests {
		result := indexNonWhiteSpace(test.input, test.in)
		if result != test.expected {
			t.Errorf("indexNonWhiteSpace(%v, %c) = %d; want %d", test.input, test.in, result, test.expected)
		}
	}
}
