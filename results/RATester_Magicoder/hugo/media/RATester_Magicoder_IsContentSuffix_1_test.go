package media

import (
	"fmt"
	"testing"
)

func TestIsContentSuffix_1(t *testing.T) {
	contentTypes := ContentTypes{
		extensionSet: map[string]bool{
			".html": true,
			".md":   true,
			".txt":  true,
		},
	}

	tests := []struct {
		name     string
		suffix   string
		expected bool
	}{
		{
			name:     "Existing suffix",
			suffix:   ".html",
			expected: true,
		},
		{
			name:     "Non-existing suffix",
			suffix:   ".doc",
			expected: false,
		},
		{
			name:     "Empty suffix",
			suffix:   "",
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

			result := contentTypes.IsContentSuffix(test.suffix)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
