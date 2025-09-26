package pageparser

import (
	"fmt"
	"testing"
)

func TestIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    []byte
		sep      []byte
		expected int
	}

	tests := []test{
		{[]byte("hello world"), []byte("world"), 6},
		{[]byte("hello world"), []byte("foo"), -1},
		{[]byte(""), []byte(""), 0},
	}

	for _, test := range tests {
		l := &pageLexer{input: test.input}
		result := l.index(test.sep)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
