package strings

import (
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
	}{
		{
			name:     "TrimSuffix_0",
			suffix:   "ing",
			s:        "string",
			expected: "str",
		},
		{
			name:     "TrimSuffix_1",
			suffix:   "g",
			s:        "string",
			expected: "strin",
		},
		{
			name:     "TrimSuffix_2",
			suffix:   "string",
			s:        "string",
			expected: "",
		},
		{
			name:     "TrimSuffix_3",
			suffix:   "not_exist",
			s:        "string",
			expected: "string",
		},
		{
			name:     "TrimSuffix_4",
			suffix:   123,
			s:        "string",
			expected: "",
		},
		{
			name:     "TrimSuffix_5",
			suffix:   "string",
			s:        123,
			expected: "",
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
			if err != nil {
				t.Errorf("TrimSuffix() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("TrimSuffix() = %v, want %v", result, tt.expected)
			}
		})
	}
}
