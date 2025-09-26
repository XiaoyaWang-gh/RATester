package modules

import (
	"fmt"
	"testing"
)

func TestIsValid_2(t *testing.T) {
	tests := []struct {
		name     string
		version  HugoVersion
		expected bool
	}{
		{
			name: "valid version",
			version: HugoVersion{
				Min: "0.100.0",
				Max: "0.102.0",
			},
			expected: true,
		},
		{
			name: "invalid version",
			version: HugoVersion{
				Min: "0.102.0",
				Max: "0.100.0",
			},
			expected: false,
		},
		{
			name: "extended version",
			version: HugoVersion{
				Min:      "0.100.0",
				Max:      "0.102.0",
				Extended: true,
			},
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

			result := test.version.IsValid()
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
