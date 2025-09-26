package related

import (
	"fmt"
	"testing"
)

func TestHasType_1(t *testing.T) {
	type test struct {
		name     string
		config   Config
		input    string
		expected bool
	}

	tests := []test{
		{
			name: "Test case 1",
			config: Config{
				Indices: IndicesConfig{
					{Type: "type1"},
					{Type: "type2"},
					{Type: "type3"},
				},
			},
			input:    "type2",
			expected: true,
		},
		{
			name: "Test case 2",
			config: Config{
				Indices: IndicesConfig{
					{Type: "type1"},
					{Type: "type2"},
					{Type: "type3"},
				},
			},
			input:    "type4",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.config.HasType(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
