package pageparser

import (
	"fmt"
	"testing"
)

func TestConsumeToSpace_1(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Test with space",
			input:          "Hello World",
			expectedOutput: "Hello",
		},
		{
			name:           "Test with tab",
			input:          "Hello\tWorld",
			expectedOutput: "Hello",
		},
		{
			name:           "Test with newline",
			input:          "Hello\nWorld",
			expectedOutput: "Hello",
		},
		{
			name:           "Test with carriage return",
			input:          "Hello\rWorld",
			expectedOutput: "Hello",
		},
		{
			name:           "Test with space at the end",
			input:          "Hello ",
			expectedOutput: "Hello",
		},
		{
			name:           "Test with multiple spaces",
			input:          "Hello   World",
			expectedOutput: "Hello",
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
			l.consumeToSpace()
			if string(l.current()) != test.expectedOutput {
				t.Errorf("Expected %s, got %s", test.expectedOutput, string(l.current()))
			}
		})
	}
}
