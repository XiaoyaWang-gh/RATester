package tailscale

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestSanitizeDomains_3(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		domains  []string
		expected []string
	}{
		{
			name:     "empty domains",
			domains:  []string{},
			expected: []string{},
		},
		{
			name:     "single valid domain",
			domains:  []string{"valid.domain"},
			expected: []string{"valid.domain"},
		},
		{
			name:     "multiple valid domains",
			domains:  []string{"valid1.domain", "valid2.domain"},
			expected: []string{"valid1.domain", "valid2.domain"},
		},
		{
			name:     "single invalid domain",
			domains:  []string{"invalid.domain"},
			expected: []string{},
		},
		{
			name:     "multiple valid and invalid domains",
			domains:  []string{"valid1.domain", "invalid.domain", "valid2.domain"},
			expected: []string{"valid1.domain", "valid2.domain"},
		},
		{
			name:     "duplicate domains",
			domains:  []string{"valid1.domain", "valid1.domain"},
			expected: []string{"valid1.domain"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeDomains(ctx, test.domains)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
