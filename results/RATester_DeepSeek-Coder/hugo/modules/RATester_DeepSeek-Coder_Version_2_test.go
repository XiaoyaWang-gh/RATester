package modules

import (
	"fmt"
	"testing"
)

func TestVersion_2(t *testing.T) {
	tests := []struct {
		name     string
		adapter  *moduleAdapter
		expected string
	}{
		{
			name: "Test with version set",
			adapter: &moduleAdapter{
				version: "1.0.0",
			},
			expected: "1.0.0",
		},
		{
			name: "Test with GoMod and version not set",
			adapter: &moduleAdapter{
				gomod: &goModule{
					Version: "1.0.0",
				},
			},
			expected: "1.0.0",
		},
		{
			name: "Test with GoMod and version set",
			adapter: &moduleAdapter{
				version: "1.0.0",
				gomod: &goModule{
					Version: "2.0.0",
				},
			},
			expected: "1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.adapter.Version()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
