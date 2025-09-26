package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestTrimSuffix_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		suffix   any
		s        any
		expected string
		err      error
	}{
		{
			name:     "TrimSuffix_Success",
			suffix:   "world",
			s:        "hello world",
			expected: "hello ",
			err:      nil,
		},
		{
			name:     "TrimSuffix_Fail",
			suffix:   "world",
			s:        123,
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

			result, err := ns.TrimSuffix(tt.suffix, tt.s)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Expected error %v, got %v", tt.err, err)
			}

			if result != tt.expected {
				t.Errorf("Expected result %v, got %v", tt.expected, result)
			}
		})
	}
}
