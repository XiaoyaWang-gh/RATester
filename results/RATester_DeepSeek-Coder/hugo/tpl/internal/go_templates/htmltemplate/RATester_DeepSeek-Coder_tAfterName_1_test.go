package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTAfterName_1(t *testing.T) {
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
				state: stateTag,
			},
			inputS: []byte("="),
			expected: context{
				state: stateBeforeValue,
			},
			expectedI: 1,
		},
		{
			name: "Test case 2",
			input: context{
				state: stateTag,
			},
			inputS: []byte("abc"),
			expected: context{
				state: stateTag,
			},
			expectedI: 0,
		},
		{
			name: "Test case 3",
			input: context{
				state: stateBeforeValue,
			},
			inputS: []byte("abc"),
			expected: context{
				state: stateBeforeValue,
			},
			expectedI: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, actualI := tAfterName(tt.input, tt.inputS)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
			if actualI != tt.expectedI {
				t.Errorf("Expected %v, but got %v", tt.expectedI, actualI)
			}
		})
	}
}
