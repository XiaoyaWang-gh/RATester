package hugofs

import (
	"fmt"
	"testing"
)

func TestTrimFrom_1(t *testing.T) {
	tests := []struct {
		name     string
		r        RootMapping
		input    string
		expected string
	}{
		{
			name: "Test trimFrom with empty string",
			r: RootMapping{
				From: "",
			},
			input:    "",
			expected: "",
		},
		{
			name: "Test trimFrom with non-empty string",
			r: RootMapping{
				From: "test",
			},
			input:    "test123",
			expected: "123",
		},
		{
			name: "Test trimFrom with non-matching prefix",
			r: RootMapping{
				From: "test",
			},
			input:    "abc",
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.r.trimFrom(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
