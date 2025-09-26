package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTJSDelimited_1(t *testing.T) {
	type test struct {
		name     string
		input    []byte
		expected context
		length   int
	}

	tests := []test{
		{
			name:  "Test case 1",
			input: []byte(`test`),
			expected: context{
				state: stateJS,
				jsCtx: jsCtxDivOp,
			},
			length: 4,
		},
		{
			name:  "Test case 2",
			input: []byte(`test\`),
			expected: context{
				state: stateError,
				err:   errorf(ErrPartialEscape, nil, 0, "unfinished escape sequence in JS string: %q", []byte(`test\`)),
			},
			length: 5,
		},
		{
			name:  "Test case 3",
			input: []byte(`test[`),
			expected: context{
				state: stateError,
				err:   errorf(ErrPartialCharset, nil, 0, "unfinished JS regexp charset: %q", []byte(`test[`)),
			},
			length: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, length := tJSDelimited(context{}, tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
			if length != tt.length {
				t.Errorf("Expected length %v, but got %v", tt.length, length)
			}
		})
	}
}
