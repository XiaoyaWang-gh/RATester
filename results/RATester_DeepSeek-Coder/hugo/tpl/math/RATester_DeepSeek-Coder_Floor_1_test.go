package math

import (
	"fmt"
	"testing"
)

func TestFloor_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected float64
		wantErr  bool
	}{
		{
			name:     "Test with positive float",
			input:    1.5,
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "Test with negative float",
			input:    -1.5,
			expected: -2,
			wantErr:  false,
		},
		{
			name:     "Test with zero",
			input:    0,
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Test with non-float value",
			input:    "string",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Floor(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Floor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Floor() = %v, want %v", got, tt.expected)
			}
		})
	}
}
