package common

import (
	"fmt"
	"testing"
)

func TestBytesToNum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    []byte
		expected int
	}

	tests := []test{
		{[]byte{1, 2, 3}, 123},
		{[]byte{4, 5, 6}, 456},
		{[]byte{7, 8, 9}, 789},
	}

	for _, test := range tests {
		result := BytesToNum(test.input)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
