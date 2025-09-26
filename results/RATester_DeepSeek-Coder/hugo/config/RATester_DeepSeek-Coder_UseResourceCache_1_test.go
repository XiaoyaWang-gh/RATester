package config

import (
	"errors"
	"fmt"
	"testing"
)

func TestUseResourceCache_1(t *testing.T) {
	tests := []struct {
		name     string
		b        BuildConfig
		err      error
		expected bool
	}{
		{
			name: "should return false when UseResourceCacheWhen is 'never'",
			b: BuildConfig{
				UseResourceCacheWhen: "never",
			},
			err:      nil,
			expected: false,
		},
		{
			name: "should return true when UseResourceCacheWhen is 'always'",
			b: BuildConfig{
				UseResourceCacheWhen: "always",
			},
			err:      nil,
			expected: true,
		},
		{
			name: "should return true when UseResourceCacheWhen is 'fallback' and err is nil",
			b: BuildConfig{
				UseResourceCacheWhen: "fallback",
			},
			err:      nil,
			expected: true,
		},
		{
			name: "should return false when UseResourceCacheWhen is 'fallback' and err is not nil",
			b: BuildConfig{
				UseResourceCacheWhen: "fallback",
			},
			err:      errors.New("some error"),
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

			result := test.b.UseResourceCache(test.err)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
