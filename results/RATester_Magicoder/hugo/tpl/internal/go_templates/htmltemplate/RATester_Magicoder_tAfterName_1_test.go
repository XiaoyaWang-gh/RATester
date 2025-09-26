package template

import (
	"fmt"
	"testing"
)

func TesttAfterName_1(t *testing.T) {
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
				state: stateTag,
			},
			input: []byte("="),
			expected: context{
				state: stateBeforeValue,
			},
			length: 1,
		},
		{
			name: "Test case 2",
			context: context{
				state: stateBeforeValue,
			},
			input: []byte("value"),
			expected: context{
				state: stateTag,
			},
			length: 5,
		},
		{
			name: "Test case 3",
			context: context{
				state: stateTag,
			},
			input: []byte("name"),
			expected: context{
				state: stateTag,
			},
			length: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			resultContext, resultLength := tAfterName(test.context, test.input)
			if resultContext.state != test.expected.state {
				t.Errorf("Expected state: %v, got: %v", test.expected.state, resultContext.state)
			}
			if resultLength != test.length {
				t.Errorf("Expected length: %v, got: %v", test.length, resultLength)
			}
		})
	}
}
