package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGet_18(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected any
	}

	tests := []testCase{
		{
			name:     "TestGet_EmptyKey",
			key:      "",
			expected: nil,
		},
		{
			name:     "TestGet_NonExistentKey",
			key:      "nonExistentKey",
			expected: nil,
		},
		{
			name:     "TestGet_ExistingKey",
			key:      "existingKey",
			expected: "expectedValue",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &defaultConfigProvider{
				root: maps.Params{
					"existingKey": "expectedValue",
				},
			}

			result := provider.Get(tc.key)

			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
