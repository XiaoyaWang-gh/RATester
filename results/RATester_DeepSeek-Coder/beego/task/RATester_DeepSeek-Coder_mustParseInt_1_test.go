package task

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMustParseInt_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected uint
		wantErr  bool
	}

	tests := []test{
		{
			name:     "Valid positive number",
			input:    "123",
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Invalid negative number",
			input:    "-123",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Non-numeric string",
			input:    "abc",
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

			if tt.wantErr {
				assert.Panics(t, func() { mustParseInt(tt.input) })
			} else {
				assert.Equal(t, tt.expected, mustParseInt(tt.input))
			}
		})
	}
}
