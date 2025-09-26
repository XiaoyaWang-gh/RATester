package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIsExist_2(t *testing.T) {
	ctx := context.Background()
	cache := &MemoryCache{
		items: map[string]*MemoryItem{
			"key1": {
				val:         "value1",
				createdTime: time.Now(),
				lifespan:    time.Hour,
			},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "Existing key",
			key:      "key1",
			expected: true,
		},
		{
			name:     "Non-existing key",
			key:      "key2",
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

			actual, err := cache.IsExist(ctx, test.key)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
