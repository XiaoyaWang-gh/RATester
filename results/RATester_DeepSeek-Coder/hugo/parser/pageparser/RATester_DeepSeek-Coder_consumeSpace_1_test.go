package pageparser

import (
	"fmt"
	"testing"
)

func TestConsumeSpace_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "consumeSpace with leading spaces",
			input:    "   hello",
			expected: "hello",
		},
		{
			name:     "consumeSpace with trailing spaces",
			input:    "hello   ",
			expected: "hello",
		},
		{
			name:     "consumeSpace with leading and trailing spaces",
			input:    "   hello   ",
			expected: "hello",
		},
		{
			name:     "consumeSpace with no spaces",
			input:    "hello",
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &pageLexer{input: []byte(test.input)}
			l.consumeSpace()
			if string(l.current()) != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, string(l.current()))
			}
		})
	}
}
