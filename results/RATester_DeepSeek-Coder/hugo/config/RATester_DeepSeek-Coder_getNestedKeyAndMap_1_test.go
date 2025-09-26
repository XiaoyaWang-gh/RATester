package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetNestedKeyAndMap_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		create   bool
		expected string
	}{
		{
			name:     "test case 1",
			key:      "key1.key2.key3",
			create:   true,
			expected: "key3",
		},
		{
			name:     "test case 2",
			key:      "key1.key2.key3",
			create:   false,
			expected: "key3",
		},
		{
			name:     "test case 3",
			key:      "key1.key2.key4",
			create:   true,
			expected: "",
		},
		{
			name:     "test case 4",
			key:      "key1.key2.key4",
			create:   false,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &defaultConfigProvider{
				root: maps.Params{
					"key1": maps.Params{
						"key2": maps.Params{
							"key3": "value3",
						},
					},
				},
			}

			actualKey, actualMap := c.getNestedKeyAndMap(tt.key, tt.create)
			if actualKey != tt.expected {
				t.Errorf("Expected key %s, but got %s", tt.expected, actualKey)
			}
			if tt.expected != "" {
				if _, ok := actualMap[tt.expected]; !ok {
					t.Errorf("Expected key %s to exist in the map", tt.expected)
				}
			}
		})
	}
}
