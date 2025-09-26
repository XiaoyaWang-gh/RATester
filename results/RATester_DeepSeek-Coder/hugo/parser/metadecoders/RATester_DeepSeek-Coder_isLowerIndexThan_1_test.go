package metadecoders

import (
	"fmt"
	"testing"
)

func TestIsLowerIndexThan_1(t *testing.T) {
	type test struct {
		name     string
		first    int
		others   []int
		expected bool
	}

	tests := []test{
		{
			name:     "first is -1",
			first:    -1,
			others:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "others are lower than first",
			first:    5,
			others:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "others are not lower than first",
			first:    5,
			others:   []int{6, 7, 8},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isLowerIndexThan(tt.first, tt.others...)
			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
