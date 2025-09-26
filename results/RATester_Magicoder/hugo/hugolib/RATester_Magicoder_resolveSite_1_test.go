package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestresolveSite_1(t *testing.T) {
	hugoSites := &HugoSites{
		Sites: []*Site{
			{
				language: &langs.Language{
					Lang: "en",
				},
			},
			{
				language: &langs.Language{
					Lang: "fr",
				},
			},
		},
	}

	tests := []struct {
		name     string
		lang     string
		expected *Site
	}{
		{
			name:     "Existing language",
			lang:     "en",
			expected: hugoSites.Sites[0],
		},
		{
			name:     "Non-existing language",
			lang:     "de",
			expected: nil,
		},
		{
			name:     "Empty language",
			lang:     "",
			expected: hugoSites.Sites[0],
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := hugoSites.resolveSite(test.lang)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
