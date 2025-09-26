package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTBlockCmt_1(t *testing.T) {
	type test struct {
		name      string
		input     context
		inputS    []byte
		expected  context
		expectedI int
	}

	tests := []test{
		{
			name: "Test case 1",
			input: context{
				state: stateJSBlockCmt,
			},
			inputS: []byte("/* block comment */"),
			expected: context{
				state: stateJS,
			},
			expectedI: 17,
		},
		{
			name: "Test case 2",
			input: context{
				state: stateJS,
			},
			inputS: []byte("/* block comment */"),
			expected: context{
				state: stateJS,
			},
			expectedI: 17,
		},
		{
			name: "Test case 3",
			input: context{
				state: stateCSSBlockCmt,
			},
			inputS: []byte("/* block comment */"),
			expected: context{
				state: stateJS,
			},
			expectedI: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, actualI := tBlockCmt(tt.input, tt.inputS)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
			if actualI != tt.expectedI {
				t.Errorf("Expected %v, but got %v", tt.expectedI, actualI)
			}
		})
	}
}
