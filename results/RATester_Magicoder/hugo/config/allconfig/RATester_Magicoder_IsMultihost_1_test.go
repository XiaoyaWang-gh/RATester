package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestIsMultihost_1(t *testing.T) {
	type testCase struct {
		name     string
		config   ConfigLanguage
		expected bool
	}

	testCases := []testCase{
		{
			name: "Single language",
			config: ConfigLanguage{
				m: &Configs{
					Languages: langs.Languages{
						{},
					},
					IsMultihost: false,
				},
			},
			expected: false,
		},
		{
			name: "Multi-language",
			config: ConfigLanguage{
				m: &Configs{
					Languages: langs.Languages{
						{},
						{},
					},
					IsMultihost: true,
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.config.IsMultihost()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
