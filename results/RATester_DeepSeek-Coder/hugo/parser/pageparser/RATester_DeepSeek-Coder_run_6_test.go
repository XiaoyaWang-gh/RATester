package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRun_6(t *testing.T) {
	type test struct {
		name     string
		input    []byte
		expected *pageLexer
	}

	tests := []test{
		{
			name:  "test case 1",
			input: []byte("test input"),
			expected: &pageLexer{
				input: []byte("test input"),
			},
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

			l := &pageLexer{
				input: tt.input,
			}

			result := l.run()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
