package modules

import (
	"fmt"
	"testing"
)

func TestHasModuleImport_1(t *testing.T) {
	testCases := []struct {
		name     string
		config   Config
		expected bool
	}{
		{
			name: "has module import",
			config: Config{
				Imports: []Import{
					{Path: "github.com/gohugoio/hugo"},
				},
			},
			expected: true,
		},
		{
			name: "no module import",
			config: Config{
				Imports: []Import{
					{Path: "github.com/gohugoio/hugo-theme-test"},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.config.hasModuleImport()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
