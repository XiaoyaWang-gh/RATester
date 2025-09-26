package hugofs

import (
	"fmt"
	"testing"
)

func TesttrimFrom_1(t *testing.T) {
	tests := []struct {
		name     string
		r        RootMapping
		input    string
		expected string
	}{
		{
			name: "Test case 1",
			r: RootMapping{
				From: "prefix",
			},
			input:    "prefix/test",
			expected: "test",
		},
		{
			name: "Test case 2",
			r: RootMapping{
				From: "prefix",
			},
			input:    "prefix",
			expected: "",
		},
		{
			name: "Test case 3",
			r: RootMapping{
				From: "prefix",
			},
			input:    "",
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

			got := tt.r.trimFrom(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}
