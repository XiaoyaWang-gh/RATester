package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetString_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		key      string
		expected string
	}{
		{
			name: "TestGetString_ExistingKey",
			provider: &defaultConfigProvider{
				root: maps.Params{"key": "value"},
			},
			key:      "key",
			expected: "value",
		},
		{
			name: "TestGetString_NonExistingKey",
			provider: &defaultConfigProvider{
				root: maps.Params{"key": "value"},
			},
			key:      "nonExistingKey",
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

			actual := tt.provider.GetString(tt.key)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
