package cache

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    interface{}
		expected bool
	}

	tests := []test{
		{input: true, expected: true},
		{input: false, expected: false},
		{input: "true", expected: true},
		{input: "false", expected: false},
		{input: "1", expected: true},
		{input: "0", expected: false},
		{input: 1, expected: true},
		{input: 0, expected: false},
		{input: "string", expected: false},
		{input: nil, expected: false},
	}

	for _, test := range tests {
		result := GetBool(test.input)
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
