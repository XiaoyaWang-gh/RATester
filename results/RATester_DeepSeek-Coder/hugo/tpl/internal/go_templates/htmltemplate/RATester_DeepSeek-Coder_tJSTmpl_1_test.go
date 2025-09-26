package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTJSTmpl_1(t *testing.T) {
	type test struct {
		name      string
		input     context
		inputS    []byte
		expected  context
		expectedN int
	}

	tests := []test{
		{
			name: "Test case 1",
			input: context{
				state: stateJS,
			},
			inputS: []byte("\\${"),
			expected: context{
				state:        stateJS,
				jsBraceDepth: []int{0, 1},
			},
			expectedN: 3,
		},
		{
			name: "Test case 2",
			input: context{
				state: stateJS,
			},
			inputS: []byte("`"),
			expected: context{
				state: stateJS,
			},
			expectedN: 1,
		},
		{
			name: "Test case 3",
			input: context{
				state: stateJS,
			},
			inputS: []byte("\\"),
			expected: context{
				state: stateError,
				err:   errorf(ErrPartialEscape, nil, 0, "unfinished escape sequence in JS string: %q", []byte("\\")),
			},
			expectedN: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, actualN := tJSTmpl(tt.input, tt.inputS)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
			if actualN != tt.expectedN {
				t.Errorf("Expected %v, but got %v", tt.expectedN, actualN)
			}
		})
	}
}
