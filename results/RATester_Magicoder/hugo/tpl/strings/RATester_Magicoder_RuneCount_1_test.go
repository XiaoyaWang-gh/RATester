package strings

import (
	"fmt"
	"testing"
)

func TestRuneCount_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected int
		wantErr  bool
	}{
		{
			name:     "valid string",
			input:    "hello",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "string with special characters",
			input:    "😀😃😄😁😆",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "non-string input",
			input:    123,
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

			got, err := ns.RuneCount(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RuneCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("RuneCount() = %v, want %v", got, tt.expected)
			}
		})
	}
}
