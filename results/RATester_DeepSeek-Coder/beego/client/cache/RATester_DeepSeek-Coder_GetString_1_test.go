package cache

import (
	"fmt"
	"testing"
)

func TestGetString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    interface{}
		expected string
	}

	tests := []test{
		{input: "test", expected: "test"},
		{input: []byte("test"), expected: "test"},
		{input: 123, expected: "123"},
		{input: nil, expected: ""},
	}

	for _, test := range tests {
		result := GetString(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
