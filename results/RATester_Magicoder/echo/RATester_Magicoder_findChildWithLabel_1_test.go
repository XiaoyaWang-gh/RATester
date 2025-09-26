package echo

import (
	"fmt"
	"testing"
)

func TestFindChildWithLabel_1(t *testing.T) {
	n := &node{
		paramChild: &node{},
		anyChild:   &node{},
	}

	tests := []struct {
		name     string
		label    byte
		expected *node
	}{
		{
			name:     "Test with param label",
			label:    paramLabel,
			expected: n.paramChild,
		},
		{
			name:     "Test with any label",
			label:    anyLabel,
			expected: n.anyChild,
		},
		{
			name:     "Test with non-existing label",
			label:    'a',
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := n.findChildWithLabel(tt.label)
			if got != tt.expected {
				t.Errorf("findChildWithLabel() = %v, want %v", got, tt.expected)
			}
		})
	}
}
