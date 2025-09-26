package template

import (
	"fmt"
	"testing"
)

func TesttBlockCmt_1(t *testing.T) {
	tests := []struct {
		name     string
		context  context
		input    []byte
		expected context
		length   int
	}{
		{
			name: "Test case 1",
			context: context{
				state: stateJSBlockCmt,
			},
			input: []byte("/* block comment */"),
			expected: context{
				state: stateJS,
			},
			length: 18,
		},
		{
			name: "Test case 2",
			context: context{
				state: stateCSSBlockCmt,
			},
			input: []byte("/* block comment */"),
			expected: context{
				state: stateCSS,
			},
			length: 18,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			resultContext, resultLength := tBlockCmt(test.context, test.input)
			if !resultContext.eq(test.expected) {
				t.Errorf("Expected context %v, but got %v", test.expected, resultContext)
			}
			if resultLength != test.length {
				t.Errorf("Expected length %d, but got %d", test.length, resultLength)
			}
		})
	}
}
