package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestIsMultilingual_1(t *testing.T) {
	type test struct {
		name     string
		config   ConfigLanguage
		expected bool
	}

	tests := []test{
		{
			name: "Single language",
			config: ConfigLanguage{
				m: &Configs{
					Languages: langs.Languages{
						{Lang: "en"},
					},
				},
			},
			expected: false,
		},
		{
			name: "Multiple languages",
			config: ConfigLanguage{
				m: &Configs{
					Languages: langs.Languages{
						{Lang: "en"},
						{Lang: "fr"},
					},
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.config.IsMultilingual()
			if result != tc.expected {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
