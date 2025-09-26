package template

import (
	"fmt"
	"testing"
)

func TestHexDecode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		input    []byte
		expected rune
	}{
		{[]byte("0"), 0},
		{[]byte("1"), 1},
		{[]byte("2"), 2},
		{[]byte("3"), 3},
		{[]byte("4"), 4},
		{[]byte("5"), 5},
		{[]byte("6"), 6},
		{[]byte("7"), 7},
		{[]byte("8"), 8},
		{[]byte("9"), 9},
		{[]byte("a"), 10},
		{[]byte("b"), 11},
		{[]byte("c"), 12},
		{[]byte("d"), 13},
		{[]byte("e"), 14},
		{[]byte("f"), 15},
		{[]byte("A"), 10},
		{[]byte("B"), 11},
		{[]byte("C"), 12},
		{[]byte("D"), 13},
		{[]byte("E"), 14},
		{[]byte("F"), 15},
	}

	for _, test := range tests {
		result := hexDecode(test.input)
		if result != test.expected {
			t.Errorf("hexDecode(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
