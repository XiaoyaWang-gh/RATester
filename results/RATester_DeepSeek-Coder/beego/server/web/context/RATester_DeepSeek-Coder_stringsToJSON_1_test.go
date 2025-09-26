package context

import (
	"fmt"
	"testing"
)

func TestStringsToJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{"Hello, world", "Hello, world"},
		{"日本語", "\\u65e5\\u672c\\u8a9e"},
		{"", ""},
	}

	for _, test := range tests {
		result := stringsToJSON(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
