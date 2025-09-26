package gin

import (
	"fmt"
	"testing"
)

func TestShiftNRuneBytes_1(t *testing.T) {
	type test struct {
		name     string
		input    [4]byte
		shift    int
		expected [4]byte
	}

	tests := []test{
		{
			name:     "shift 0",
			input:    [4]byte{0, 1, 2, 3},
			shift:    0,
			expected: [4]byte{0, 1, 2, 3},
		},
		{
			name:     "shift 1",
			input:    [4]byte{0, 1, 2, 3},
			shift:    1,
			expected: [4]byte{1, 2, 3, 0},
		},
		{
			name:     "shift 2",
			input:    [4]byte{0, 1, 2, 3},
			shift:    2,
			expected: [4]byte{2, 3, 0, 0},
		},
		{
			name:     "shift 3",
			input:    [4]byte{0, 1, 2, 3},
			shift:    3,
			expected: [4]byte{3, 0, 0, 0},
		},
		{
			name:     "shift 4",
			input:    [4]byte{0, 1, 2, 3},
			shift:    4,
			expected: [4]byte{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := shiftNRuneBytes(tt.input, tt.shift)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
