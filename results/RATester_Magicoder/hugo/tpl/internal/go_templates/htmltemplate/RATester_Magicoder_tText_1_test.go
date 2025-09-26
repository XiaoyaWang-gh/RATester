package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttText_1(t *testing.T) {
	tests := []struct {
		name        string
		input       context
		inputS      []byte
		expected    context
		expectedLen int
	}{
		{
			name: "Test case 1",
			input: context{
				state: stateHTMLCmt,
			},
			inputS: []byte("<!-- HTML comment -->"),
			expected: context{
				state: stateHTMLCmt,
			},
			expectedLen: 21,
		},
		{
			name: "Test case 2",
			input: context{
				state: stateTag,
			},
			inputS: []byte("<tag>"),
			expected: context{
				state: stateTag,
			},
			expectedLen: 5,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := tText(tt.input, tt.inputS)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("tText() got = %v, want %v", got, tt.expected)
			}
			if got1 != tt.expectedLen {
				t.Errorf("tText() got1 = %v, want %v", got1, tt.expectedLen)
			}
		})
	}
}
