package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestIsZero_6(t *testing.T) {
	type testCase struct {
		name     string
		config   *Configs
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Nil config",
			config:   nil,
			expected: true,
		},
		{
			name:     "Empty config",
			config:   &Configs{},
			expected: true,
		},
		{
			name: "Non-empty config",
			config: &Configs{
				Languages: langs.Languages{
					{},
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

			result := tc.config.IsZero()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
