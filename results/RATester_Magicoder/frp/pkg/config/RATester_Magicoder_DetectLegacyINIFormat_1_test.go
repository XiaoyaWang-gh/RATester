package config

import (
	"fmt"
	"testing"
)

func TestDetectLegacyINIFormat_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected bool
	}{
		{
			name:     "Legacy INI format",
			content:  []byte("[common]\nkey=value"),
			expected: true,
		},
		{
			name:     "Non-legacy INI format",
			content:  []byte("key=value"),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := DetectLegacyINIFormat(test.content)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
