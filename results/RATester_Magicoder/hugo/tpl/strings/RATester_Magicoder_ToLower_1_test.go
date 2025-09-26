package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestToLower_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name:     "Test with string",
			input:    "HELLO",
			expected: "hello",
			err:      nil,
		},
		{
			name:     "Test with int",
			input:    123,
			expected: "",
			err:      errors.New("unable to cast 123 of type int to string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.ToLower(tt.input)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("ToLower() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.expected {
				t.Errorf("ToLower() = %v, want %v", got, tt.expected)
			}
		})
	}
}
