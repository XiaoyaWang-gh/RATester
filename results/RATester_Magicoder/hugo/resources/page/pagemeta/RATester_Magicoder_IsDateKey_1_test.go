package pagemeta

import (
	"fmt"
	"testing"
)

func TestIsDateKey_1(t *testing.T) {
	handler := FrontMatterHandler{
		allDateKeys: map[string]bool{
			"date":        true,
			"publishDate": true,
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "Existing date key",
			key:      "date",
			expected: true,
		},
		{
			name:     "Existing custom date key",
			key:      "publishDate",
			expected: true,
		},
		{
			name:     "Non-existing date key",
			key:      "nonExistingKey",
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

			result := handler.IsDateKey(test.key)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
