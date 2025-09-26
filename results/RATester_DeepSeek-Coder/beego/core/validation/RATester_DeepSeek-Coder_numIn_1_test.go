package validation

import (
	"fmt"
	"testing"
)

func TestNumIn_1(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		err      error
	}{
		{
			name:     "existingFunc",
			expected: 2,
			err:      nil,
		},
		{
			name:     "nonExistingFunc",
			expected: 0,
			err:      fmt.Errorf("doesn't exists %s valid function", "nonExistingFunc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			num, err := numIn(tt.name)
			if num != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, num)
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
