package media

import (
	"fmt"
	"testing"
)

func TestGetBestMatch_1(t *testing.T) {
	tests := []struct {
		name     string
		types    Types
		input    string
		expected Type
		found    bool
	}{
		{
			name: "GetByType",
			types: Types{
				{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			},
			input:    "rss",
			expected: Type{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			found:    true,
		},
		{
			name: "GetBySubType",
			types: Types{
				{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			},
			input:    "application",
			expected: Type{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			found:    true,
		},
		{
			name: "GetFirstBySuffix",
			types: Types{
				{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			},
			input:    "xml",
			expected: Type{Type: "application/rss+xml", MainType: "application", SubType: "rss"},
			found:    true,
		},
		{
			name:     "Not Found",
			types:    Types{},
			input:    "notfound",
			expected: Type{},
			found:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, found := test.types.GetBestMatch(test.input)
			if result != test.expected || found != test.found {
				t.Errorf("Expected (%v, %v), got (%v, %v)", test.expected, test.found, result, found)
			}
		})
	}
}
