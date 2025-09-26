package inflect

import (
	"fmt"
	"testing"
)

func TestPluralize_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected string
		wantErr  bool
	}{
		{
			name:     "Test with a string",
			input:    "apple",
			expected: "apples",
			wantErr:  false,
		},
		{
			name:     "Test with a number",
			input:    1,
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Test with a nil",
			input:    nil,
			expected: "",
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

			got, err := ns.Pluralize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Pluralize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Namespace.Pluralize() = %v, want %v", got, tt.expected)
			}
		})
	}
}
