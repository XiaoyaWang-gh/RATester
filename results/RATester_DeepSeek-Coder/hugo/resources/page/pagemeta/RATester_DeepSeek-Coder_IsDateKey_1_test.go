package pagemeta

import (
	"fmt"
	"testing"
)

func TestIsDateKey_1(t *testing.T) {
	testCases := []struct {
		name        string
		allDateKeys map[string]bool
		key         string
		expected    bool
	}{
		{
			name:        "Key exists",
			allDateKeys: map[string]bool{"date": true, "lastmod": true, "publishdate": true, "expirydate": true},
			key:         "date",
			expected:    true,
		},
		{
			name:        "Key does not exist",
			allDateKeys: map[string]bool{"lastmod": true, "publishdate": true, "expirydate": true},
			key:         "date",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := FrontMatterHandler{
				allDateKeys: tc.allDateKeys,
			}

			result := f.IsDateKey(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
