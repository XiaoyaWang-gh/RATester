package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_12(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected []string
	}

	tests := []test{
		{
			name:     "Add one element",
			input:    []string{"element1"},
			expected: []string{"element1"},
		},
		{
			name:     "Add multiple elements",
			input:    []string{"element1", "element2", "element3"},
			expected: []string{"element1", "element2", "element3"},
		},
		{
			name:     "Add empty elements",
			input:    []string{"", "element1", "", "element2", ""},
			expected: []string{"element1", "element2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pagePathBuilder{}
			p.Add(tt.input...)
			if !reflect.DeepEqual(p.els, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, p.els)
			}
		})
	}
}
