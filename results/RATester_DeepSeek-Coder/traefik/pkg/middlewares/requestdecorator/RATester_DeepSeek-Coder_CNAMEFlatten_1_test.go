package requestdecorator

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestCNAMEFlatten_1(t *testing.T) {
	ctx := context.Background()
	hr := &Resolver{
		CnameFlattening: true,
		ResolvConfig:    "8.8.8.8",
		ResolvDepth:     3,
		cache:           cache.New(30*time.Minute, 5*time.Minute),
	}

	testCases := []struct {
		name     string
		host     string
		expected string
	}{
		{
			name:     "test case 1",
			host:     "google.com",
			expected: "google.com",
		},
		{
			name:     "test case 2",
			host:     "github.com",
			expected: "github.com",
		},
		{
			name:     "test case 3",
			host:     "invalid.host",
			expected: "invalid.host",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := hr.CNAMEFlatten(ctx, tc.host)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
