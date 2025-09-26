package config

import (
	"fmt"
	"testing"
)

func TestSplitEnvVar_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			name  string
			value string
		}
	}{
		{
			name:  "Test case 1",
			input: "NAME=VALUE",
			expected: struct {
				name  string
				value string
			}{
				name:  "NAME",
				value: "VALUE",
			},
		},
		{
			name:  "Test case 2",
			input: "NAME=",
			expected: struct {
				name  string
				value string
			}{
				name:  "NAME",
				value: "",
			},
		},
		{
			name:  "Test case 3",
			input: "=VALUE",
			expected: struct {
				name  string
				value string
			}{
				name:  "",
				value: "VALUE",
			},
		},
		{
			name:  "Test case 4",
			input: "NAME",
			expected: struct {
				name  string
				value string
			}{
				name:  "NAME",
				value: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			name, value := SplitEnvVar(test.input)
			if name != test.expected.name || value != test.expected.value {
				t.Errorf("Expected (%s, %s), got (%s, %s)", test.expected.name, test.expected.value, name, value)
			}
		})
	}
}
